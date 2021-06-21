package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatReadMsg(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) error {

	// 接收参数
	var msg vo.ChatReadMsgParams

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
	userId := claim["UserId"].(string)

	for _, id := range msg.Ids {
		res, err := db.DB.Exec("UPDATE chat_msg SET is_read = 1 WHERE id = ? AND to_id = ?", id, userId)
		if err != nil {
			sugar.Log.Error("UPDATE chat_msg is_read failed.Err is ", err)
			return err
		}

		_, err = res.RowsAffected()
		if err != nil {
			sugar.Log.Error("UPDATE chat_msg is_read failed.Err is ", err)
			return err
		}
	}

	return nil
}
