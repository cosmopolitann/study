package mvc

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//递归删除文件夹下面所有的文件或者文件夹
var delArray []string

func del(db *Sql, parent_id string, userId string) {
	//如果有子文件夹，则：
	if parent_id != "" {
		log.Println("parent_id = ", parent_id)
		var f File
		var f1 []File
		rows, _ := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0) FROM cloud_file where parent_id=? and user_id=?", parent_id, userId)
		// 释放锁
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder)
			if err != nil {
				log.Println("find err is ", err)
			}
			if f.Id != "" {
				f1 = append(f1, f)
			}
		}

		for i := 0; i < len(f1); i++ {
			delArray = append(delArray, f1[i].Id)
			if f1[i].IsFolder == 1 {
				del(db, f1[i].Id, userId)
			}
		}
	}
	log.Println("All delete ids : ", delArray)
}
func Delete(db *Sql, value string) error {
	sugar.Log.Info("~~~~ Start   delete file   ~~~~ ")
	var d vo.CloudDeleteParams
	err := json.Unmarshal([]byte(value), &d)
	//marshal params.
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info(" Marshal params :=", d)
	//verify token.
	claim, b := jwt.JwtVeriyToken(d.Token)
	if !b {
		return err
	}
	sugar.Log.Info("claim := ", claim)
	// if it's a folder, delete it recursively.
	// if it's a file,delete it directly.
	// query for all ids to delete
	for _, v := range d.Ids {
		rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0) from cloud_file where id=?", v)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return errors.New("查询下载列表信息失败")
		}

		var dl File
		for rows.Next() {
			err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder)
			if err != nil {
				sugar.Log.Error("Query scan data is failed.The err is ", err)
				return err
			}
		}
		// 释放锁
		rows.Close()
		if dl.IsFolder == 1 {
			del(db, dl.Id, claim["id"].(string))
		}
		delArray = append(delArray, string(v))
	}
	//delete all ids.
	sugar.Log.Info("All ids := ", delArray)
	// Open the transaction.
	tx, err := db.DB.Begin()
	if err != nil {
		return errors.New(" Delete file or folder is failed. ")
	}
	for _, v := range delArray {
		stmt, err := db.DB.Prepare("delete from cloud_file where id=?")
		if err != nil {
			sugar.Log.Error(" Delete file is failed.Err:", err)
			tx.Rollback()
			return errors.New(" Delete file is failed. ")
		}
		res, err := stmt.Exec(v)
		if err != nil {
			sugar.Log.Error("Delete file is failed.Err::", err)
			tx.Rollback()
			return errors.New(" Delete file is failed. ")
		}
		log.Println(res)
	}
	err = tx.Commit()
	if err != nil {
		sugar.Log.Info("Delete file is failed.Err:", err)
		return errors.New(" Delete file is failed. ")
	}
	sugar.Log.Info(" Array ids : ", delArray)
	return nil

}
