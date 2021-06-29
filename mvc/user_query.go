package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
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
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') from sys_user where id=?", claim["UserId"])
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl, err
	}
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

func UserUpdate(db *Sql, value string) (e error) {
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
	userid := claim["UserId"]
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
	sugar.Log.Info("~~~~~   update user  is Successful ~~~~~~")
	return nil
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
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') from sys_user where id=?", userlist.UserId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl, err
	}
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
