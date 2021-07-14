package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func CloudFileRename(db *Sql, value string) error {

	var art vo.CloudFileRenameParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//查询是否存在记录
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(art.Token)
	if !b {
		return errors.New("token 失效")
	}
	//userid:=claim["id"].(string)
	sugar.Log.Info("claim := ", claim)
	var dl File
	//查询数据
	//stmt, err := db.DB.Query("select * from cloud_file where file_name=? and is_folder=? and parent_id",)
	rows, err := db.DB.Query("select a.id from cloud_file as a where file_name=? and is_folder=? and parent_id", art.Rename, art.IsFolder, art.ParentId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return errors.New("查询下载列表信息失败")
	}

	// 释放锁
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&dl.Id)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	sugar.Log.Info("查询到要重命名的文件是", dl)

	if dl.Id != "" {
		return errors.New("文件名称已经存在")
	} else {
		//更新
		stmt, err := db.DB.Prepare("update cloud_file set file_name=? where id=?")
		if err != nil {
			sugar.Log.Error("Update  data is failed.The err is ", err)
			return err
		}
		res, err := stmt.Exec(art.Rename, art.Id)
		if err != nil {
			sugar.Log.Error("Update  is failed.The err is ", err)
			return err
		}
		c, _ := res.RowsAffected()
		if c == 0 {
			return errors.New("更新失败")
		}
	}

	return nil

}
