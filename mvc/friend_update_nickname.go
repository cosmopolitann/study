package mvc

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"
)

func FriendUpdateNickname(db *Sql, value string) error {

	// 用户在线检查参数
	var msg vo.FriendUpdateNicknameParams

	sugar.Log.Debug("Request Param:", value)

	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", msg)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userId := claim["id"].(string)
	friendId := msg.FriendId
	friendNickname := msg.Nickname

	var oldNickname string
	err = db.DB.QueryRow("select friend_nickname from user_friend where user_id = ? and friend_id = ?", userId, friendId).Scan(&oldNickname)
	if err != nil && err != sql.ErrNoRows {
		sugar.Log.Error("query user_friend failed, err:", err)
		return err
	}

	if err == sql.ErrNoRows {
		// insert
		_, err := db.DB.Exec("insert into user_friend (user_id, friend_id, friend_nickname) values (?, ?, ?)", userId, friendId, friendNickname)
		if err != nil {
			sugar.Log.Error("insert into user_friend failed, err:", err)
			return err
		}
	} else {
		// update
		_, err := db.DB.Exec("update user_friend set friend_nickname = ? where user_id = ? and friend_id = ?", friendNickname, userId, friendId)
		if err != nil {
			sugar.Log.Error("update user_friend nickname failed, err:", err)
			return err
		}
	}

	return nil
}
