package article

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
)

func TestAddArticle(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	// 	value := `{
	//    "id": "325707698052770816",
	//    "accesstory": "QMabcdefghijk96",
	//    "text":"正文开始内容",
	//    "accesstory_type": 0,
	//    "tag": "标签33",
	//    "ptime": "1623955566",
	//   "play_num": 0,
	//    "share_num": 0,
	//    "title": "title65",
	//    "user_id": "323733228975432704",
	//    "thumbnail": "thumbnail5",
	//    "file_name": "",
	//    "file_size": ""
	// }`
	// 	var art vo.ArticleAddParams
	// 	err = json.Unmarshal([]byte(value), &art)
	// 	if err != nil {
	// 		sugar.Log.Error("Marshal is failed.Err is ", err)
	// 	}
	// 	sugar.Log.Info("Marshal data is  ", art)
	// id := utils.SnowId()
	// t1:=time.Now().Unix()
	// stmt, err := d.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	stmt, err := d.Prepare("INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size) values ('asdfa123sdf','13414','123',2,'123','213',1312,12312,0,'fgh','123','nijk',1)")

	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
	}
	// sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec()
	//sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t1, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
	}
	l, _ := res.RowsAffected()
	if l == 0 {
	}
	// resp := (value)
	// fmt.Println("这是返回的数据 =", resp)
	// 添加的sql 语句 写入 文件中
	f1, err1 := os.OpenFile("/Users/apple/winter/offline/update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
	if err1 != nil {
		fmt.Println("创建失败")
	}
	fmt.Println("  ----- 本地 local 文件 存在  ----")
	//拼接sql 语句
	sql := "INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size) values ('asdfa23','13414','123',2,'123','213',1312,12312,0,'fgh','123','nijk',1)\n"

	_, err = f1.WriteString(sql)
	if err != nil {
		fmt.Println(" 写入 local 文件 错误：", err)
	}
	fmt.Println(" 写入 local 文件 成功 1", err)

	sql1 := "INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size) values ('asdfa78','13414','123',2,'123','213',1312,12312,0,'fgh','123','nijk',1)\n"

	_, err = f1.WriteString(sql1)
	if err != nil {
		fmt.Println(" 写入 local 文件 错误：", err)
	}
	fmt.Println(" 写入 local 文件 成功 2", err)

}
func Testdb2(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
