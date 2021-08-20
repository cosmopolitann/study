package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"testing"
)

func TestAddFolder(t *testing.T) {
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

	var f1 = vo.CloudAddFolderParams{
		Id:       "123",
		FileName: "_月亮湾",
		ParentId: "3",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjI5NjU3OTMwfQ.4P3tY5xUkKdUCHSlXliRnDDdrHK_cQNtskZgl6kTbWY",
	}

	b1, e := json.Marshal(f1)
	fmt.Println(e)
	fmt.Println(string(b1))
	resp := ss.AddFolder(string(b1))
	fmt.Println("这是返回的数据 =", resp)

}
