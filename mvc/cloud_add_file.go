package mvc

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
)

func AddFile(db *Sql, value string) (string, error) {
	//add file
	var f vo.CloudAddFileParams
	err := json.Unmarshal([]byte(value), &f)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("解析数据:", f)
	if err != nil {
		sugar.Log.Error("Decode is failed.", err)
		return "", errors.New("decode is failed")
	}

	//token verify
	claim, b := jwt.JwtVeriyToken(f.Token)
	if !b {
		return "", err
	}
	userId := claim["UserId"]
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO cloud_file (id,user_id,file_name,parent_id,ptime,file_cid,file_size,file_type,is_folder,thumbnail) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into cloud_file table is failed.", err)
		return "", err
	}
	sid := strconv.FormatInt(id, 10)
	res, err := stmt.Exec(sid, userId, f.FileName, f.ParentId, t, f.FileCid, f.FileSize, f.FileType, 0, f.Thumbnail)
	if err != nil {
		sugar.Log.Error("Insert into file  is Failed.", err)
		return "", err
	}
	sugar.Log.Info("Insert into file  is successful.")
	l, _ := res.RowsAffected()
	if l == 0 {
		return "", err
	}
	return sid, nil

}
func FindOneFileIsExist(db *Sql, ff map[string]interface{}, f File) (int64, error) {
	//查询数据
	rows, _ := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') FROM cloud_file where file_name=? and parent_id=?", ff["FileName"], ff["ParentId"])
	for rows.Next() {
		err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder, &f.Thumbnail)
		if err != nil {
			return 0, err
		}
	}
	if f.Id != "" {
		return 1, nil
	}
	return 0, nil
}
