package user

import (
	"github.com/cosmopolitann/clouddb/sugar"

	"github.com/cosmopolitann/clouddb/mvc"

	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestUserRegister(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "../../tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	//插入数据
	var fi = mvc.File{
		Id:       "1",
		UserId:   "408217533556985856",
		FileName: "红楼梦",
		ParentId: "0",
		FileCid:  "Qmcid",
		FileSize: 100,
		FileType: 11,
		IsFolder: 0,
		Ptime:    1232131,
	}
	b1, e := json.Marshal(fi)
	fmt.Println(e)
	fmt.Println(b1)

	//这里 改成 穿 json 字符串，字段 要改成更新之后的数据。

	//{"id":"4324","peerId":"124","name":"20","phone":1,"sex":"1","nickName":"nick"}
	value := `{"id":"43243421","peerId":"Q1w213e12332211","name":"20","phone":"12233456","sex":"1","nickName":"nick","img":"123"}`
	//resp:= ss.UserAdd(string(b1)

	resp := ss.UserRegister(nil,value)
	fmt.Println("这是返回的数据 =", resp)
}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}

func TestExportUser(t *testing.T) {
	
}