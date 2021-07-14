package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func DownloadList(db *Sql, value string) (data []DownLoad, e error) {
	sugar.Log.Info("~~~ Start   DownloadList  ~~~")
	var d []DownLoad
	var tl vo.TransferListParams
	//marshal params.
	err := json.Unmarshal([]byte(value), &tl)
	if err != nil {
		return d, err
	}

	//check token.
	claim, b := jwt.JwtVeriyToken("")
	if !b {
		return d, errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(down_path,'null'),IFNULL(file_type,0),IFNULL(transfer_type,0),IFNULL(upload_parent_id,0),IFNULL(upload_file_id,0) from cloud_transfer where user_id=?", claim["id"].(string))
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return d, errors.New("查询下载列表信息失败")
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		var dl DownLoad
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.DownPath, &dl.FileType, &dl.TransferType, &dl.UploadParentId, &dl.UploadFileId)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return d, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		d = append(d, dl)
	}
	sugar.Log.Info("Query all data is ", d)
	return d, nil
}
