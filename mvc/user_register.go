package mvc

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func AddUser(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string, path string) (vo.UserLoginRespParams, error) {
	sugar.Log.Info(" ~~~ Start Add User  ~~~ ")
	sugar.Log.Info(" ----  Path :", path)
	//user string ==> user struct
	//Add sys_user
	//create snow id
	var resp vo.UserLoginRespParams
	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		return resp, err
	}
	sugar.Log.Info("params ：= ", user)
	//create snowId
	id := utils.SnowId()
	//create now time
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user (id,peer_id,name,phone,sex,ptime,utime,nickname,img,role) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed:", err.Error())
		return resp, err
	}
	sid := strconv.FormatInt(id, 10)
	user.Phone = sid //手机号注册不用了,phone字段直接用id来填,兼容老版本
	user.Sex = 0
	user.NickName = "dragon" + sid[len(sid)-5:]

	res, err := stmt.Exec(sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img, user.Role)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed:", err.Error())
		return resp, err
	}
	c, _ := res.RowsAffected()
	sugar.Log.Info("~~~~~   Insert into sys_user data is Successful ~~~~~~", c)
	//生成 token
	// 手机号
	//token,err:=jwt.GenerateToken(user.Phone,60)

	resp.Token, _ = jwt.GenerateToken(sid, user.PeerId, user.Name, user.Phone, user.NickName, user.Img, user.Role, user.Sex, user.Ptime, user.Utime, -1)
	resp.UserInfo = GetUser(db, sid)
	// publish msg
	topic := "/db-online-sync"
	sugar.Log.Info("Publish Topic: ", "/db-online-sync")
	sugar.Log.Info("Publish recieve: ", value)
	ctx := context.Background()
	//Topic join.
	tp, ok := TopicJoin.Load(topic)
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return resp, err
		}
		TopicJoin.Store(topic, tp)
	}
	sugar.Log.Info("--- Start publish msg ---")
	sugar.Log.Info("--- first step ---")

	sugar.Log.Info("publish content :", value)
	// query data about publish msg.
	var dl vo.RespSysUser
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null'),IFNULL(role,'2') from sys_user where id=?", sid)
	if err != nil {
		sugar.Log.Error("AddUser Query data is failed.Err is ", err)
		return resp, err
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName, &dl.Img, &dl.Role)
		if err != nil {
			sugar.Log.Error("AddUser Query scan data is failed.The err is ", err)
			return resp, err
		}
		sugar.Log.Info(" AddUser Query a entire data is ", dl)
	}
	//the first step.
	var s3 UserAd
	s3.Type = "receiveUserRegister"
	s3.Data = dl
	s3.FromId = ipfsNode.Identity.String()
	//marshal UserAd.
	//the second step
	sugar.Log.Info("--- second step ---")

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Error("Publish msg is failed.Err:", err)
		return resp, err
	}
	sugar.Log.Info("Frwarding information:=", string(jsonBytes))
	sugar.Log.Info("Local PeerId :=", ipfsNode.Identity.String())
	//the  third  step .
	sugar.Log.Info("--- third step ---")

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
		return resp, err
	}
	sugar.Log.Info("~~~~  Publish msg is successful.   ~~~~  ")
	// write sql to update file.
	//because update file need upload to ipfs.
	sugar.Log.Info("--- Write sql to update file. ---")
	f1, err1 := os.OpenFile(path+"update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
	if err1 != nil {
		sugar.Log.Error("Create update file is failed.Err: ", err1)
	}
	sugar.Log.Info("----- Local update is already exist.  ----")
	//
	rand.Seed(time.Now().UnixNano())

	//step2：获取随机数
	num4 := rand.Intn(1000000) + 0
	s1 := strconv.Itoa(num4)
	nickname := "dragon" + s1
	// sprintf sql
	sql := fmt.Sprintf("INSERT INTO sys_user (id,peer_id,name,phone,sex,ptime,utime,nickname,img,role) values('%s','%s','%s','%s',%d,%d,%d,'%s','%s','%s')\n", sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t, nickname, user.Img, user.Role)
	_, err = f1.WriteString(sql)
	if err != nil {
		sugar.Log.Error(" Write update file is failed.Err: ", err)
	}
	sugar.Log.Info(" Sql : ", sql)
	sugar.Log.Info("~~~~  Write update is successful.~~~~ ")
	return resp, nil
}

type UserAd struct {
	Type string `json:"type"`

	Data vo.RespSysUser `json:"data"`

	FromId string `json:"from"`
}

func FindIsExistUser(db *Sql, user SysUser) (int64, error) {
	var s SysUser
	sugar.Log.Info("start sys_user is exist local user info.")
	sugar.Log.Info("user info is ", user.Phone)
	sugar.Log.Info("user info is ", user)

	rows, _ := db.DB.Query("SELECT id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') FROM sys_user where phone=?", user.Phone)

	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName, &s.Img)
		if err != nil {
			sugar.Log.Error(" query is failed. ", err)

			return 0, err
		}
		sugar.Log.Info(" user info is ", s)
	}
	//is exist
	sugar.Log.Info(" FindOne data is ", s.Id)

	if s.Id != "" {
		return 1, nil
	} else {
		return 0, nil
	}

}
func AddUserTest(db *Sql, value string) (vo.UserLoginRespParams, error) {
	sugar.Log.Info(" ~~~ Start Add User  ~~~ ")
	//user string ==> user struct
	//Add sys_user
	//create snow id
	var resp vo.UserLoginRespParams
	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		return resp, err
	}
	sugar.Log.Info("params ：= ", user)
	//create snowId
	id := utils.SnowId()
	//create now time
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user (id,peer_id,name,phone,sex,ptime,utime,nickname,img) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed:", err.Error())
		return resp, err
	}
	sid := strconv.FormatInt(id, 10)
	user.Phone = sid //手机号注册不用了,phone字段直接用id来填,兼容老版本
	user.Sex = 0
	user.NickName = "dragon" + sid[len(sid)-5:]
	res, err := stmt.Exec(sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed:", err.Error())
		return resp, err
	}
	c, _ := res.RowsAffected()
	sugar.Log.Info("~~~~~   Insert into sys_user data is Successful ~~~~~~", c)
	//生成 token
	// 手机号
	//token,err:=jwt.GenerateToken(user.Phone,60)

	resp.Token, _ = jwt.GenerateToken(user.Id, user.PeerId, user.Name, user.Phone, user.NickName, user.Img, "2", user.Sex, user.Ptime, user.Utime, -1)
	resp.UserInfo = GetUser(db, sid)

	return resp, nil
}
