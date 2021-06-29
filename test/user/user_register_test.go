package user

import (
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"

	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestUserRegister(t *testing.T) {
	//a,err1 := jwt.GenerateToken("123",-1)
	//fmt.Println("token is:",a)
	//time.Sleep(time.Second*3)
	//tmp,err1 := jwt.ParseToken(a)
	//if err1 != nil {
	//	fmt.Println("token解析出错:",err1.Error())
	//	return
	//}
	//fmt.Println("解析结果:",*tmp)
	//return
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
	//{"id":"4324","peerId":"124","name":"20","phone":1,"sex":"1","nickName":"nick"}

	// value := `{"id":"43243421","peerId":"Q1w213e1233221","name":"20","phone":"12233456","sex":"1","nickName":"nick","img":"123"}`
	// //resp:= ss.UserAdd(string(b1)
	// resp := ss.UserRegister(value)
	// fmt.Println("这是返回的数据 =", resp)

	// value := `{"id":"43243421","peerId":"Q1w2112312323221111","name":"20","phone":"12233456","sex":"1","nickName":"nick","img":"123"}`
	//resp:= ss.UserAdd(string(b1)

	// resp := ss.UserRegister(nil, value)
	// fmt.Println("这是返回的数据 =", resp) //这里 改成 穿 json 字符串，字段 要改成更新之后的数据。

}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}

func TestExportUser(t *testing.T) {

}
