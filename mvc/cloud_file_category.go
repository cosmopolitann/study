package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//根据 文件类型 进行分类

func CloudFileCategory(db *Sql, value string) (data []File, e error) {
	var list vo.FileCategoryParams
	var arrfile []File
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", list)
	//
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(list.Token)
	if !b {
		return arrfile, errors.New("token 失效")
	}
	userid := claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	// 排序
	var or string
	if list.Order == "" {
		or = "ptime"
	}
	if list.Order == "time" {
		or = "ptime"
	}
	if list.Order == "name" {
		or = "file_name"
	}
	if list.Order == "type" {
		or = "file_type"

	}
	if list.Order == "size" {
		or = "file_size"
	}
	sugar.Log.Info("排序方式:", or)
	// 查询
	rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where file_type=? and user_id=? order by ?", list.FileType, userid, or)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, err
	}
	// 释放锁
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
