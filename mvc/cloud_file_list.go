package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//查询文件列表

func CloudFileList(db *Sql, value string) (data []File, e error) {
	var list vo.CloudFindListParams
	var arrfile []File

	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", list)
	//验证 token 是否 满足条件
	claim, b := jwt.JwtVeriyToken(list.Token)
	userId := claim["UserId"]
	sugar.Log.Info("userId := ", userId)

	if !b {
		return arrfile, errors.New("token 验证失败")
	}
	sugar.Log.Info("claim := ", claim)

	// 查询
	rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where parent_id=? and is_folder=? and user_id=?", "0", 0, "123")
	// rows, err := db.DB.Query("select * from cloud_file where parent_id=? and is_folder=? and user_id=?", "0", 0, "409330202166956032")

	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	defer rows.Close()
	for rows.Next() {
		var dl File
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder, &dl.Thumbnail)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("Query all data is ", arrfile)
	return arrfile, nil
}
