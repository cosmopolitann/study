package mvc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
)

//用户查询信息

func UserQuery(db *Sql, value string) (data SysUser, e error) {
	var dl SysUser
	var userlist vo.UserListParams
	err := json.Unmarshal([]byte(value), &userlist)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", userlist)
	//verify token
	claim, b := jwt.JwtVeriyToken(userlist.Token)
	if !b {
		return dl, err
	}
	sugar.Log.Info("claim := ", claim)
	//query
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') from sys_user where id=?", claim["id"])
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl, err
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName, &dl.Img)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	sugar.Log.Info("~~~~~   Delete user into  is Successful ~~~~~~")
	return dl, nil
}

// 更新用户信息

func UserUpdate(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) (e error) {
	var userlist vo.UserUpdateParams
	err := json.Unmarshal([]byte(value), &userlist)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is ", userlist)
	//check token
	claim, b := jwt.JwtVeriyToken(userlist.Token)
	if !b {
		return errors.New(" Token is invalid. ")
	}
	userid := claim["id"]
	sugar.Log.Info("claim:= ", claim)
	//user info.
	sugar.Log.Info("User Info := ", userlist)
	//update data.
	stmt, err := db.DB.Prepare("update sys_user set sex=?,nickname=?,img=? where id=?")
	if err != nil {
		sugar.Log.Error("Prepare is failed.Err:= ", err)
		return err
	}
	res, err := stmt.Exec(userlist.Sex, userlist.NickName, userlist.Img, userid)
	if err != nil {
		sugar.Log.Error("Exec is failed.Err:= ", err)
		return err
	}
	affect, _ := res.RowsAffected()
	if affect == 0 {
		return errors.New(" update user info  is failed. ")
	}
	var dl vo.RespSysUser

	//查询用户信息
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null'),IFNULL(role,'2') from sys_user where id=?", userid)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName, &dl.Img, &dl.Role)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}

	//
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
			return err
		}
		TopicJoin.Store(topic, tp)
	}
	var s3 UserAd
	s3.Type = "receiveUserUpdate"
	s3.Data = dl
	s3.FromId = ipfsNode.Identity.String()
	//marshal UserAd.
	//the second step
	sugar.Log.Info("--- second step ---")

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Error("Publish msg is failed.Err:", err)
		return err
	}
	sugar.Log.Info("Frwarding information:=", string(jsonBytes))
	sugar.Log.Info("Local PeerId :=", ipfsNode.Identity.String())
	//the  third  step .
	sugar.Log.Info("--- third step ---")

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
		return err
	}
	sugar.Log.Info("~~~~  Publish msg is successful.   ~~~~  ")
	//
	sugar.Log.Info("~~~~~   update user  is Successful ~~~~~~")
	return nil
}

type UserUpdateInfo struct {
	Type string `json:"type"`

	Data vo.RespSysUser `json:"data"`

	FromId string `json:"from"`
}

// 查询 对方 用户 信息

//用户查询信息

func OtherUserQuery(db *Sql, value string) (data SysUser, e error) {
	var dl SysUser
	var userlist vo.OtherUserInfoParams
	err := json.Unmarshal([]byte(value), &userlist)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", userlist)
	sugar.Log.Info(" UserId := ", userlist.UserId)
	//query
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null'),IFNULL(role,'2') from sys_user where id=?", userlist.UserId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl, err
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName, &dl.Img)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	sugar.Log.Info("~~~~~   Delete user into  is Successful ~~~~~~")
	return dl, nil
}
