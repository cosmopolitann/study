package mvc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

var M []string

var MM [][]string

//CopyFile
func CopyFile(db *Sql, value string) error {
	//copy file or  copy dir
	var s File
	var cFile vo.CopyFileParams
	err := json.Unmarshal([]byte(value), &cFile)
	if err != nil {
		sugar.Log.Error("解析错误:", err)
		return err
	}

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(cFile.Token)
	if !b {
		return errors.New("token 失效")
	}
	userid := claim["UserId"].(string)
	// for _, v := range cFile.Ids {
	// 	rows, err := db.DB.Query("SELECT b.id,IFNULL(b.user_id,'null'),IFNULL(b.file_name,'null'),IFNULL(b.parent_id,0),IFNULL(b.ptime,0),IFNULL(b.file_cid,'null'),IFNULL(b.file_size,0),IFNULL(b.file_type,0),IFNULL(b.is_folder,0),IFNULL(b.thumbnail,'null') from cloud_file as b WHERE (b.file_name,b.user_id,b.is_folder) in (SELECT a.file_name,a.user_id,a.is_folder from cloud_file as a where a.id=?) and b.parent_id=?", v, cFile.ParentId)
	// 	if err != nil {
	// 		sugar.Log.Error("Select cloud_file is failed.", err)
	// 		return err
	// 	}
	// 	for rows.Next() {
	// 		err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid, &s.FileSize, &s.FileType, &s.IsFolder, &s.Thumbnail)
	// 		if err != nil {
	// 			sugar.Log.Error("Scan is failed.", err)
	// 			return err
	// 		}
	// 	}
	// 	if s.Id != "" {
	// 		return errors.New("文件已经存在")
	// 	}
	// 	if s.Id == "" {
	// 		//0  文件  1 文件夹
	// 		rows1, err1 := db.DB.Query("SELECT b.id,IFNULL(b.user_id,'null'),IFNULL(b.file_name,'null'),IFNULL(b.parent_id,0),IFNULL(b.ptime,0),IFNULL(b.file_cid,'null'),IFNULL(b.file_size,0),IFNULL(b.file_type,0),IFNULL(b.is_folder,0),IFNULL(b.thumbnail,'null') from cloud_file as b WHERE b.id=?", v)
	// 		if err1 != nil {
	// 			sugar.Log.Error("Select cloud_file is failed.", err1)

	// 			return errors.New("查询文件失败")
	// 		}
	// 		for rows1.Next() {

	// 			err := rows1.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid, &s.FileSize, &s.FileType, &s.IsFolder, &s.Thumbnail)

	// 			if err != nil {
	// 				sugar.Log.Error("Scan is failed.", err)
	// 				return err
	// 			}
	// 		}
	// 		sugar.Log.Infof("query data is s:= ", s)

	// 		id := utils.SnowId()
	// 		// t := time.Now().Format("2006-01-02 15:04:05")
	// 		t := time.Now().Unix()
	// 		stmt, err := db.DB.Prepare("INSERT INTO cloud_file (id,user_id,file_name,parent_id,ptime,file_cid,file_size,file_type,is_folder,thumbnail) values(?,?,?,?,?,?,?,?,?,?)")
	// 		if err != nil {
	// 			sugar.Log.Error("Insert into cloud_file table is failed.", err)
	// 			return err
	// 		}
	// 		sid := strconv.FormatInt(id, 10)
	// 		res, errt := stmt.Exec(sid, userid, s.FileName, cFile.ParentId, t, s.FileCid, s.FileSize, s.FileType, s.IsFolder, s.Thumbnail)
	// 		if errt != nil {
	// 			return errors.New("插入文件失败")
	// 		}
	// 		c, _ := res.RowsAffected()
	// 		if c == 0 {
	// 			return errors.New("插入文件失败")
	// 		}
	// 	}
	// }
	for _, v := range cFile.Ids {
		rows, err := db.DB.Query("SELECT b.id,IFNULL(b.user_id,'null'),IFNULL(b.file_name,'null'),IFNULL(b.parent_id,0),IFNULL(b.ptime,0),IFNULL(b.file_cid,'null'),IFNULL(b.file_size,0),IFNULL(b.file_type,0),IFNULL(b.is_folder,0),IFNULL(b.thumbnail,'null') from cloud_file as b WHERE (b.file_name,b.user_id,b.is_folder) in (SELECT a.file_name,a.user_id,a.is_folder from cloud_file as a where a.id=?) and b.parent_id=?", v, cFile.ParentId)
		if err != nil {
			sugar.Log.Error("Select cloud_file is failed.", err)
			return err
		}
		for rows.Next() {
			err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid, &s.FileSize, &s.FileType, &s.IsFolder, &s.Thumbnail)
			if err != nil {
				sugar.Log.Error("Scan is failed.", err)
				return err
			}
		}
		if s.Id != "" {
			return errors.New("文件已经存在")
		}

	}
	//都满足
	M = []string{}
	for _, v := range cFile.Ids {
		M = append(M, v)
	}
	MM = append(MM, M)
	fmt.Println(" 第一次的MM ===", MM)
	M = []string{}

	for _, v := range cFile.Ids {
		// 查询 是否是文件夹 在传进去。

		rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where id=?", v)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)

		}
		for rows.Next() {
			var dl File
			err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder, &dl.Thumbnail)
			if err != nil {
				sugar.Log.Error("Query scan data is failed.The err is ", err)
			}
			if dl.IsFolder == 1 {
				cc(db.DB, userid, v)
			}

		}
	}
	sugar.Log.Info("-----------------------M=", M)

	sugar.Log.Info("--------最终的 ---------------MM=", MM)
	//插入数据

	return nil
}

//

// func  (db *sql.DB, parent_id string, userId string) error {
// for _, v := range MM {
// 	for _, v1 := range v {
// 		rows, err := db.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where id=?", v1)
// 		if err != nil {
// 			sugar.Log.Error("Query data is failed.Err is ", err)

// 		}
// 		for rows.Next() {
// 			var dl File
// 			err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder, &dl.Thumbnail)
// 			if err != nil {
// 				sugar.Log.Error("Query scan data is failed.The err is ", err)
// 			}
// 		}

// 		id := utils.SnowId()
// 		t := time.Now().Unix()
// 		stmt, err := db.Prepare("INSERT INTO cloud_file (id,user_id,file_name,parent_id,ptime,file_cid,file_size,file_type,is_folder,thumbnail) values(?,?,?,?,?,?,?,?,?,?)")
// 		if err != nil {
// 			sugar.Log.Error("Insert into cloud_file table is failed.", err)
// 			return err
// 		}
// 		sid := strconv.FormatInt(id, 10)
// 		res, err := stmt.Exec(sid, userId, f.FileName, f.ParentId, t, f.FileCid, f.FileSize, f.FileType, 0, f.Thumbnail)
// 		if err != nil {
// 			sugar.Log.Error("Insert into file  is Failed.", err)
// 			return "", err
// 		}
// 		sugar.Log.Info("Insert into file  is successful.")
// 		l, _ := res.RowsAffected()
// 		if l == 0 {
// 			return "", err
// 		}
// 	}
// }

// }

func cc(d *sql.DB, user_id, id string) {
	rows, err := d.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where user_id=? and parent_id=?", user_id, id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)

	}
	for rows.Next() {
		var dl File
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder, &dl.Thumbnail)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
		}
		if dl.Id != "" {

			M = append(M, dl.Id)
		}

	}
	MM = append(MM, M)

	for _, v := range M {
		rows, err := d.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where id=?", v)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)

		}
		var dl File
		for rows.Next() {
			err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder, &dl.Thumbnail)
			if err != nil {
				sugar.Log.Error("Query scan data is failed.The err is ", err)
			}
		}
		if dl.IsFolder == 1 {
			M = []string{}
			cc(d, user_id, dl.Id)
		}
	}

	// MM = append(MM, M)
	// M = []string{}
	sugar.Log.Info("-----------------------M=", M)

	//
	sugar.Log.Info("-----------------------MM=", MM)

}

//插入
