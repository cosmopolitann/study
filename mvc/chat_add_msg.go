package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"time"
)

func AddChatMsg(db *Sql, value string) error {

	var msg vo.ChatAddMsgParams
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", msg)
	//
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userid := claim["id"].(string)
	id := utils.SnowId()
	//t := time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO chat_msg values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into chat_msg table is failed.", err)
		return err
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, msg.ContentType, msg.Content, userid, msg.ToId, t, msg.IsWithdraw, msg.IsRead, msg.RecordId)
	if err != nil {
		sugar.Log.Error("Insert into chat_msg  is Failed.", err)
		return err
	}
	sugar.Log.Info("Insert into chat_msg  is successful.")
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New("插入chat_msg数据失败")
	}
	return nil

}
