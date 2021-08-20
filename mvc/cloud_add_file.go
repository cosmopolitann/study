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
	userId := claim["id"]
	id := utils.SnowId()
	t := time.Now().Unix()
	//查询 是否 有相同名字的 文件
	c, err, snowid := FindFileSameName(db, f)
	if err != nil {
		return "", err

	}
	if c == 0 {
		stmt, err := db.DB.Prepare("INSERT INTO cloud_file (id,user_id,file_name,parent_id,ptime,file_cid,file_size,file_type,is_folder,thumbnail,width,height,duration) values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			sugar.Log.Error("Insert into cloud_file table is failed.", err)
			return "", err
		}
		sid := strconv.FormatInt(id, 10)
		res, err := stmt.Exec(sid, userId, f.FileName, f.ParentId, t, f.FileCid, f.FileSize, f.FileType, 0, f.Thumbnail, f.Width, f.Height, f.Duration)
		if err != nil {
			sugar.Log.Error("Insert into file  is Failed.", err)
			return "", err
		}
		sugar.Log.Info("Insert into file  is successful.")
		l, _ := res.RowsAffected()
		if l == 0 {
			return "", err
		}
		snowid = sid
	} else if c == 1 {
		stmt, err := db.DB.Prepare("INSERT INTO cloud_file (id,user_id,file_name,parent_id,ptime,file_cid,file_size,file_type,is_folder,thumbnail,width,height,duration) values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			sugar.Log.Error("Insert into cloud_file table is failed.", err)
			return "", err
		}
		sid := strconv.FormatInt(id, 10)
		//
		ti := strconv.FormatInt(t, 10)
		res, err := stmt.Exec(sid, userId, f.FileName+ti, f.ParentId, t, f.FileCid, f.FileSize, f.FileType, 0, f.Thumbnail)
		if err != nil {
			sugar.Log.Error("Insert into file  is Failed.", err)
			return "", err
		}
		sugar.Log.Info("Insert into file  is successful.")
		l, _ := res.RowsAffected()
		if l == 0 {
			return "", err
		}
		snowid = sid
	} else if c == 2 {
		//

	}
	return snowid, nil

}
func FindOneFileIsExist(db *Sql, ff map[string]interface{}, f File) (int64, error) {
	//查询数据
	rows, _ := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') FROM cloud_file where file_name=? and parent_id=?", ff["FileName"], ff["ParentId"])

	// 释放锁
	defer rows.Close()

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

func FindFileSameName(db *Sql, p vo.CloudAddFileParams) (int64, error, string) {
	//查询数据
	f := File{}
	rows, _ := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') FROM cloud_file where file_name=? and parent_id=?", p.FileName, p.ParentId)

	// 释放锁
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder, &f.Thumbnail)
		if err != nil {
			return 3, err, ""
		}
	}
	if f.Id != "" {
		if f.FileCid != p.FileCid {
			//插入 加时间戳
			return 1, nil, ""
		} else {
			return 2, nil, f.Id
		}
	}
	return 0, nil, ""
}
