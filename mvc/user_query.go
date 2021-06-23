package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func UserQuery(db *Sql, value string) (data SysUser, e error) {
	var dl SysUser
	var userlist vo.UserListParams
	err := json.Unmarshal([]byte(value), &userlist)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", userlist)
	//校验 token 是否 满足
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

func UserUpdate(db *Sql, value string) (e error) {
	var userlist vo.UserUpdateParams
	err := json.Unmarshal([]byte(value), &userlist)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is ", userlist)
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(userlist.Token)
	if !b {
		return errors.New("token验证失败")
	}
	userid := claim["UserId"]
	sugar.Log.Info("claim:= ", claim)
	//更新 用户 信息
	sugar.Log.Info("用户 信息:= ", userlist)
	// 判断 逻辑

	//更新数据
	stmt, err := db.DB.Prepare("update sys_user set sex=?,nickname=?,img=? where id=?")
	if err != nil {

		return err
	}
	res, err := stmt.Exec(userlist.Sex, userlist.NickName,userlist.Img, userid)
	if err != nil {

		return err
	}
	affect, _ := res.RowsAffected()

	if affect == 0 {

		return errors.New("更新失败")
	}

	sugar.Log.Info("~~~~~   update user  is Successful ~~~~~~")
	return nil
}
