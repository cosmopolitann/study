package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//删除消息

func ChatMsgDel(db *Sql, value string) error {
	var msgdel vo.ChatMsgDelParams
	err := json.Unmarshal([]byte(value), &msgdel)
	if err != nil {
		return err
	}
	sugar.Log.Info("删除的id:", msgdel.Id)
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msgdel.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	//userid:=claim["id"].(string)
	stmt, err := db.DB.Prepare("delete from chat_msg where id=?")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(msgdel.Id)
	if err != nil {
		sugar.Log.Error("delete data from chat_msg  table is failed.", err)
		return err
	}
	c, _ := res.RowsAffected()
	if c == 0 {
		sugar.Log.Error("delete data from chat_msg  table is failed.", err)
		return err
	}
	sugar.Log.Info("delete data from chat_msg  table is failed")
	return nil

}
