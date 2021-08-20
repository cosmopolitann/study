package mvc

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//MoveFile
func MoveFile(db *Sql, value string) error {
	sugar.Log.Info(" ~~~~  Start   MoveFile  ~~~~~ ")
	//move file or  copy dir
	var mvFile vo.MoveFileParams
	err := json.Unmarshal([]byte(value), &mvFile)
	if err != nil {
		sugar.Log.Info("Unmarshal 参数值 ： ", err)
		return err
	}
	//step 1  find user info is exist.
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(mvFile.Token)
	if !b {
		return errors.New(" Token is invaild.")
	}
	sugar.Log.Info("解析token 参数值 ： ", claim)
	//userid:=claim["UserId"].(string)
	for _, v := range mvFile.Ids {

		log.Println("这是要移动的文件id：", v)
		rows, _ := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null'),IFNULL(width,''),IFNULL(height,''),IFNULL(duration,0) from cloud_file as b WHERE (b.file_name,b.user_id,b.is_folder) in (SELECT a.file_name,a.user_id,a.is_folder from cloud_file as a WHERE a.id=?) and b.parent_id=?;", v, mvFile.ParentId)
		var s File

		for rows.Next() {
			err := rows.Scan(&s.Id, &s.UserId, &s.FileName, &s.ParentId, &s.Ptime, &s.FileCid, &s.FileSize, &s.FileType, &s.IsFolder, &s.Thumbnail, &s.Width, &s.Height, &s.Duration)

			if err != nil {
				fmt.Println("query err is ", err)
				return err
			}
		}
		// 释放锁
		rows.Close()

		log.Println(" 这是移动 查找出来的结果  move file  =", s)

		if s.Id != "" {
			if s.IsFolder == 1 {
				return errors.New("文件夹已经存在")
			} else {
				return errors.New("文件已经存在")
			}
		}
		if s.Id == "" {
			log.Println(" 这是移动 查找出来的结果  move file  =", s)
			log.Println(" 文件不存在 =", s)
			log.Println(" movde file  id  =", s.Id)
			//0  文件  1 文件夹
			//	if s.IsFolder==0{
			log.Println(" 更新 文件 父  id 信息 =", s)

			//stmt, err := mvc.DB.Prepare("INSERT INTO cloud_file values(?,?,?,?,?,?,?,?,?,?)")
			stmt, err := db.DB.Prepare("UPDATE cloud_file set parent_id=? where id=?")

			//update userinfo set username=? where uid=?
			fmt.Println(" 这是 需要 更新的 id == ", v)
			fmt.Println(" 这是 需要 更新的 父 id  == ", mvFile.ParentId)
			res, err := stmt.Exec(mvFile.ParentId, v)
			if err != nil {
				sugar.Log.Error("Update cloud_file table is failed.", err)
				return err
			}
			c, _ := res.RowsAffected()
			if c == 0 {
				return errors.New("update is failed.")
			}
		}
	}
	sugar.Log.Info(" ~~~~  Start   MoveFile  End ~~~~~ ")
	return nil
}
