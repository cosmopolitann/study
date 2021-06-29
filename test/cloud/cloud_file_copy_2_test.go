package cloud

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

func TestCopy2(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}

	//插入数据
	
	MM = [][]string{}
	// if dl.IsFolder == 1 {
	// 	copy(d, dl.Id, "409330202166956032")
	// }
	//
	M = []string{}
	M = append(M, "414849335038054400")
	MM = append(MM, M)
	M = []string{}
	cc(d, "409330202166956032", "414849335038054400")
	fmt.Println("--------------")

	fmt.Println("最终的 MM =", MM)
	//插入数据

	for _, v := range MM {
		for k, v := range v {
			fmt.Println("k=", k)
			fmt.Println("v=", v)

		}
	}

}

func cc(d *sql.DB, user_id, id string) {
	rows, err := d.Query("select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where user_id=? and parent_id=?", user_id, id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)

	}
	for rows.Next() {
		var dl File1
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
		var dl File1
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

var M []string

var MM [][]string
var delArray []string

// var delArray1[]string

func copy(db *sql.DB, parent_id string, userId string) {
	//如果有子文件夹，则：
	if parent_id != "" {
		log.Println("parent_id = ", parent_id)
		var f File1

		var f1 []File1
		rows, _ := db.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'mk') FROM cloud_file where parent_id=? and user_id=?", parent_id, userId)
		for rows.Next() {
			err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder, &f.Thumbnail)
			if err != nil {
				log.Println("find err is ", err)
			}
			if f.Id != "" {
				f1 = append(f1, f)
				log.Println("--- ----- f1 : ", f1)
			}
			log.Println("--- ----- f1 =================: ", f1)
			for _, v := range f1 {
				M = append(M, v.Id)
				log.Println("--- ----- M : ", M)

			}
			MM = append(MM, M)
			log.Println("--- ----- MM : ", MM)

		}

		for i := 0; i < len(f1); i++ {
			delArray = append(delArray, f1[i].Id)
			if f1[i].IsFolder == 1 {
				copy(db, f1[i].Id, userId)
			}
			// for _, v := range f1 {

			// 	M = append(M, v.Id)
			// 	MM = append(MM, M)
			// }
		}
		//插入数据

	}

	log.Println("All delete ids : ", delArray)

}

type File1 struct {
	Id        string `json:"id"`        //id
	UserId    string `json:"userId"`    //用户userid
	FileName  string `json:"fileName"`  //文件名字
	ParentId  string `json:"parentId"`  //父id
	FileCid   string `json:"fileCid"`   //文件cid
	FileSize  int64  `json:"fileSize"`  //文件大小
	FileType  int64  `json:"fileType"`  //文件类型
	IsFolder  int64  `json:"isFolder"`  //是否是文件or 文件夹  0文件 1文件夹
	Ptime     int64  `json:"ptime"`     //时间
	Thumbnail string `json:"thumbnail"` //缩略图
}
