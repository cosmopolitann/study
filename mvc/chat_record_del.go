package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatRecordDel(db *Sql, value string) error {

	var msg vo.ChatRecordDelParams
	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		return err
	}

	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)

	res, err := db.DB.Exec("DELETE FROM chat_record WHERE id = ?", msg.Id)
	if err != nil {
		sugar.Log.Error("delete chat_record data is failed.", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		sugar.Log.Error("delete chat_record data is failed2.", err)
		return err
	}
	sugar.Log.Info("delete record is successful.")

	return nil

}
