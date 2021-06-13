package user

import (
	"github.com/cosmopolitann/clouddb/sugar"

	"github.com/cosmopolitann/clouddb/mvc"

	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"time"
)

func TestUserUpdate(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	fmt.Println(" Ping is failed,err:=",e)
	ss:= Testdb(d)
	//插入数据
	var fi = mvc.File{
		Id:         "1",
		UserId:     "408217533556985856",
		FileName:   "红楼梦",
		ParentId:   "0",
		FileCid:    "Qmcid",
		FileSize:   100,
		FileType:   11,
		IsFolder:   0,
		Ptime:      time.Time{},
	}
	b1, e := json.Marshal(fi)
	fmt.Println(e)
	fmt.Println(b1)

	//这里 改成 穿 json 字符串，字段 要改成更新之后的数据。

	//{"id":"4324","peerId":"124","name":"20","phone":1,"sex":"1","nickName":"nick"}
	value:=`{"id":"410199041234702336","peerId":"555555","name":"莉莉娅23","phone":"1888","sex":"0","nickName":"lillsf","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI"}`
	//resp:= ss.UserAdd(string(b1)

	resp:= ss.UserUpdate(value)

	fmt.Println("这是返回的数据 =",resp)

}