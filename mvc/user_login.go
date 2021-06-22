package mvc

import (
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func UserLogin(db *Sql, value string) (vo.UserLoginRespParams, error) {
	//解析传进来的参数信息
	var resp vo.UserLoginRespParams
	//改用token登陆
	claim, b := jwt.JwtVeriyToken(value)
	if !b {
		return resp,errors.New("token 失效")
	}
	userid := claim["UserId"].(string)
	user := GetUser(db, userid)
	if user.Id == "" {
		return resp, errors.New("请先注册用户")
	}
	/*
		//解析传进来的参数信息
			var resp vo.UserLoginRespParams

			var userLogin vo.UserLoginParams
			err := json.Unmarshal([]byte(value), &userLogin)
		if err != nil {
			sugar.Log.Error("Marshal is failed.Err is ", err)
		}
		sugar.Log.Info("解析登录的参数数据:", userLogin)
		//先查询数据库是否已经注册   如果 未注册 请注册  注册了 生成token
		//查询数据库
		c, err, user := FindIsExistLoginUser(db, userLogin.Phone)
		if err != nil {
			return resp, err
		}
		if c == 0 {
			return resp, errors.New("请先注册用户")
		}
	*/
	////生成 token,暂时用手机号 后面会改成唯一识别的秘钥。
	token, err := jwt.GenerateToken(user.Id, -1)
	if err != nil {
		return resp, errors.New("生成token失败，请重新登录")
	}
	//1
	resp.Token = token
	resp.UserInfo = user
	sugar.Log.Info("登录返回的信息:", resp)
	return resp, nil
}

func GetUser(db *Sql, userid string) vo.RespSysUser {
	var s vo.RespSysUser
	rows, _ := db.DB.Query("SELECT id,peer_id,name,phone,sex,ptime,utime,nickname,img FROM sys_user as a where id=?", userid)
	for rows.Next() {
		err := rows.Scan(&s.Id, &s.PeerId, &s.Name, &s.Phone, &s.Sex, &s.Ptime, &s.Utime, &s.NickName, &s.Img)
		if err != nil {
			sugar.Log.Error("查找用户表失败,原因:", err)
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
