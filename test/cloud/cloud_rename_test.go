package cloud

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestRenameFile(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Desktop/xiaolong.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	//插入数据
	value := `{
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNjIwMzU1NjI5MTM1NDYyNCIsInBlZXJJZCI6IlFtMTIzIiwibmFtZSI6Im5pY2siLCJwaG9uZSI6IjE4MSIsInNleCI6MCwibmlja25hbWUiOiJzZGYiLCJpbWciOiJodHRwIiwicHRpbWUiOjEsInV0aW1lIjoxLCJleHAiOjE2MjY1NTE3ODd9.CELlZoQgrRyElp8NwSt4QGq-XrIB0ZNJJA1LnI85Mgc",
    "rename":"4.png",
    "id":"4123",
"isFolder":0,
"parentId":"0"
}`
	//b1, e := json.Marshal(fi)
	//fmt.Println(ss)
	//fmt.Println(b1)
	resp := ss.FileRename(string(value))
	fmt.Println("这是返回的数据 =", resp)

}
