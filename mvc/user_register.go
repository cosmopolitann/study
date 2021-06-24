package mvc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func AddUser(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string, path string) (vo.UserLoginRespParams, error) {
	sugar.Log.Info(" ----  Path :", path)

	//user string ==> user struct
	//Add sys_user
	//create snow id
	var resp vo.UserLoginRespParams
	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {

	}
	sugar.Log.Info("params ：= ", user)
	/** 手机号注册不用了,改用直接注册
	l, e := FindIsExistUser(db, user)
	if e != nil {
		sugar.Log.Error("FindIsExistUser info is Failed.")
	}
	// l > 0 user is exist.
	sugar.Log.Error("-----------1")

	if l > 0 {
		sugar.Log.Error("user is exist.")
		return errors.New("user is exist.")
	}
	*/
	//inExist insert into sys_user.
	sugar.Log.Info("-----------用户 信息 -------", user)

	id := utils.SnowId()
	//create now time
	//t:=time.Now().Format("2006-01-02 15:04:05")
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
	resp.Token, _ = jwt.GenerateToken(sid, -1)
	resp.UserInfo = GetUser(db, sid)
	//return resp,nil
	//=====
	// publish msg
	topic := "/db-online-sync"
	sugar.Log.Info("发布主题:", "/db-online-sync")
	sugar.Log.Info("发布消息:", value)
	//判断是否弃用
	var tp *pubsub.Topic
	var ok bool
	ctx := context.Background()
	if tp, ok = Topicmp["/db-online-sync"]; ok == false {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			return resp, err
		}
		Topicmp[topic] = tp

	}
	sugar.Log.Info("--- 开始 发布的消息 ---")

	sugar.Log.Info("发布的消息:", value)
	//=====
	//查询数据
	var dl vo.RespSysUser

	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') from sys_user where id=?", sid)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return resp, err
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName, &dl.Img)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return resp, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}

	//================================

	//第一步
	var s3 UserAd
	s3.Type = "receiveUserRegister"
	s3.Data = dl
	s3.FromId = ipfsNode.Identity.String()
	//

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Info("--- 开始 发布的消息 ---")
		return resp, err
	}
	sugar.Log.Info("--- 解析后的数据 返回给 转接服务器 ---", string(jsonBytes))
	sugar.Log.Info("--- 这是 节点的id 信息 ---", ipfsNode.Identity.String())

	//====

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return resp, err
	}
	sugar.Log.Info("---  发布的消息  完成  ---")
	// 写入文件
	sugar.Log.Info("---  sql语句 写入文件  ---")
	//
	f1, err1 := os.OpenFile(path+"update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
	if err1 != nil {
		sugar.Log.Error("创建失败update :", err1)
	}

	sugar.Log.Info("----- 本地 local 文件 存在  ----")

	// 拼接字符串 sql 语句
	sql := fmt.Sprintf("INSERT INTO sys_user (id,peer_id,name,phone,sex,ptime,utime,nickname,img) values('%s','%s','%s','%s',%d,%d,%d,'%s','%s')\n", sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img)

	_, err = f1.WriteString(sql)
	if err != nil {
		sugar.Log.Error("-----  写入 local 文件 错误：  ----", err)
	}
	sugar.Log.Info("-----  写入 local 文件 成功 ----", err)
	//
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
