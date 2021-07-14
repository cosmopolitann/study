package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//查询文件列表

func TransferList(db *Sql, value string) (data []TransferDownLoadParams, e error) {
	sugar.Log.Info(" ~~~~  Start   TransferList ~~~~~ ")
	var list vo.TransferListParams
	var arrfile []TransferDownLoadParams
	//marshal params.
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return arrfile, err
	}
	sugar.Log.Info("Marshal data is  ", list)
	//check token
	claim, b := jwt.JwtVeriyToken(list.Token)
	if !b {
		return arrfile, errors.New(" Token is invaild. ")
	}
	sugar.Log.Info("claim := ", claim)

	rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(down_path,'null'),IFNULL(file_type,0),IFNULL(transfer_type,0),IFNULL(upload_parent_id,0),IFNULL(upload_file_id,0) from cloud_transfer where user_id=?", claim["id"].(string))
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		var dl TransferDownLoadParams
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.DownPath, &dl.FileType, &dl.TransferType, &dl.UploadParentId, &dl.UploadFileId)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)
	sugar.Log.Info(" ~~~~  Start   TransferList  End ~~~~~ ")
	return arrfile, nil

}
