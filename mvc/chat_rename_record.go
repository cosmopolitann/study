package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatRenameRecord(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) error {

	// 接收参数
	var msg vo.ChatRenameRecordParams

	sugar.Log.Info("Request Param:", value)

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
	// userId := claim["UserId"].(string)

	res, err := db.DB.Exec("UPDATE chat_record SET name = ? WHERE id = ?", msg.Name, msg.Id)
	if err != nil {
		sugar.Log.Error("UPDATE chat_record name failed.Err is ", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		sugar.Log.Error("UPDATE chat_record name failed.Err is ", err)
	}

	return nil
}
