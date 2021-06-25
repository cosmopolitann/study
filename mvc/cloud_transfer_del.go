package mvc

import (
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

// 删除传输信息

func TransferDel(db *Sql, value string) error {
	sugar.Log.Info(" ~~~~  Start   TransferDel ~~~~~ ")
	var dFile vo.TransferDelParams
	//marshal params.
	err := json.Unmarshal([]byte(value), &dFile)
	if err != nil {
		sugar.Log.Error(" Marshal is failed.Err:", err)
		return err
	}
	//check token.
	claim, b := jwt.JwtVeriyToken(dFile.Token)
	if !b {
		return err
	}
	sugar.Log.Info("claim := ", claim)
	// open the transaction.
	for _, v := range dFile.Ids {
		tx, _ := db.DB.Begin()

		//todo
		stmt, err := db.DB.Prepare("delete from cloud_transfer where id=?")
		if err != nil {
			return err
		}
		res, err := stmt.Exec(v)
		if err != nil {
			sugar.Log.Error("Insert into cloud_file table is failed.", err)
			//rowback
			tx.Rollback()
			return err
		}
		c, _ := res.RowsAffected()
		if c == 0 {
			tx.Rollback()
		}
		fmt.Println(res)
		tx.Commit()
	}
	sugar.Log.Info("Insert into file  is successful.")
	sugar.Log.Info(" ~~~~  Start   TransferDel   End ~~~~~ ")

	return nil

}
