package mvc

import (
	"errors"
	"fmt"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//用户登录

func UserLogin(db *Sql, value string) (vo.UserLoginRespParams, error) {
	sugar.Log.Info("~~~~ Start User Login  ~~~")

	var resp vo.UserLoginRespParams
	//改用token登陆
	claim, b := jwt.JwtVeriyToken(value)
	if !b {
		return resp, errors.New(" Token is invaild. ")
	}
	userid := claim["UserId"].(string)
	user := GetUser(db, userid)
	if user.Id == "" {
		return resp, errors.New("请先注册用户")
	}
	token, err := jwt.GenerateToken(user.Id, -1)
	if err != nil {
		return resp, errors.New("生成token失败，请重新登录")
	}

	resp.Token = token
	resp.UserInfo = user
	sugar.Log.Info("Login resp msg:=", resp)
	sugar.Log.Info("~~~~ Start User Login   End ~~~")
	return resp, nil
}

func GetUser(db *Sql, userid string) vo.RespSysUser {
	fmt.Println("获得用户,id:", userid)
	var s vo.RespSysUser
	rows, err := db.DB.Query("SELECT id,peer_id,name,phone,sex,ptime,utime,nickname,img FROM sys_user as a where id = ? ", userid)
	if err != nil {
		sugar.Log.Error("查找用户表失败,sql错误:", err.Error())
		return s
	}
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName, &s.Img)
		if err != nil {
			sugar.Log.Error("查找用户表失败,原因:", err.Error())
			return s
		}
		sugar.Log.Info("用户信息:", s)
	}
	//is exist
	sugar.Log.Info("查找到的用户信息: ", s.Id)
	return s
}

func FindIsExistLoginUser(db *Sql, data string) (int64, error, vo.RespSysUser) {
	var s vo.RespSysUser
	sugar.Log.Info("用户信息是", data)
	rows, _ := db.DB.Query("SELECT id,peer_id,name,phone,sex,ptime,utime,nickname,img FROM sys_user as a where phone=?", data)
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName, &s.Img)
		if err != nil {
			sugar.Log.Error("查找用户表失败,原因:", err)
			return 0, err, s
		}
		sugar.Log.Info("用户信息:", s)
	}
	//is exist
	sugar.Log.Info("查找到的用户信息: ", s.Id)

	if s.Id != "" {
		//说明 有 这个用户
		return 1, nil, s
	} else {
		//说明 没有 这个用户
		return 0, nil, s
	}

}
