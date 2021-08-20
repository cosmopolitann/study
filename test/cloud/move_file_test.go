package cloud

import (
	"database/sql"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

//CopyFile
func TestMoveFile(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Desktop/xiaolong1.db")
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
"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjI5NjU3OTMwfQ.4P3tY5xUkKdUCHSlXliRnDDdrHK_cQNtskZgl6kTbWY",
    "parentId":"0",
    "ids":["435120994852540416"]
}`
	//b1, e := json.Marshal(fi)
	//fmt.Println(ss)
	//fmt.Println(b1)
	resp := ss.MoveFile(string(value))
	fmt.Println("这是返回的数据 =", resp)

}
