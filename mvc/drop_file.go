package mvc

import (
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/sugar"
)

func DeleteOneFile(db *Sql, value string) error {

	var dFile DeleteParams
	err := json.Unmarshal([]byte(value), &dFile)
	if err != nil {
		return err
	}
	for _, v := range dFile.DropFile {
		tx, _ := db.DB.Begin()

		stmt, err := db.DB.Prepare("delete from sys_user where id=?")
		checkErr(err)
		res, err := stmt.Exec(v.Id)
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
	return nil
}
