package chat

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatRecordDel(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/xiaolong.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","id":"90000001"}
`
	resp := ss.ChatRecordDel(value)
	t.Log("获取返回的数据 :=  ", resp)

}
