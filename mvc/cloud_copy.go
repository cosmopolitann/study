package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"time"
)

//CopyFile
func CopyFile(db *Sql, value string) error {
	//copy file or  copy dir
	var s File
	var cFile vo.CopyFileParams
	err := json.Unmarshal([]byte(value), &cFile)
	if err != nil {
		sugar.Log.Error("解析错误:",err)
		return err
	}

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(cFile.Token)
	if !b {
		return errors.New("token 失效")
	}
	userid := claim["UserId"].(string)
	for _, v := range cFile.Ids {
		rows, err := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0) from cloud_file as b WHERE (b.file_name,b.user_id,b.is_folder) in (SELECT a.file_name,a.user_id,a.is_folder from cloud_file as a where a.id=?) and b.parent_id=?", v, cFile.ParentId)
		if err != nil {
			sugar.Log.Error("Select cloud_file is failed.", err)
			return err
		}
		for rows.Next() {
			err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid, &s.FileSize, &s.FileType, &s.IsFolder)
			if err != nil {
				sugar.Log.Error("Scan is failed.", err)
				return err
			}
		}
		if s.Id != "" {
			return errors.New("文件已经存在")
		}
		if s.Id == "" {
			//0  文件  1 文件夹
			for rows.Next() {
				err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid, &s.FileSize, &s.FileType, &s.IsFolder)

				if err != nil {
					sugar.Log.Error("Scan is failed.", err)
					return err
				}
			}
			id := utils.SnowId()
			t := time.Now().Format("2006-01-02 15:04:05")
			stmt, err := db.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?)")
			if err != nil {
				sugar.Log.Error("Insert into cloud_file table is failed.", err)
				return err
			}
			sid := strconv.FormatInt(id, 10)
			res, err := stmt.Exec(sid, userid, s.FileName, cFile.ParentId, t, s.FileCid, s.FileSize, s.FileType, s.IsFolder)
			c, _ := res.RowsAffected()
			if c == 0 {
				return errors.New("插入文件失败")
			}
		}
	}
	return nil
}
