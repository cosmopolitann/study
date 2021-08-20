package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestAddFile(t *testing.T) {
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
	var fi = vo.CloudAddFileParams{
		Id:       "411580511585046528",
		FileName: "我爱成都2",
		ParentId: "0",
		FileCid:  "Qm123",
		FileSize: 100,
		FileType: 0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjQxNDIwNzExNDIxNTQyODA5MCIsInBlZXJJZCI6IiIsIm5hbWUiOiLmtYvor5UiLCJwaG9uZSI6IiIsInNleCI6MCwibmlja25hbWUiOiLmtYvor5UiLCJpbWciOiIiLCJwdGltZSI6MTYyNzQ0NDAwOCwidXRpbWUiOjE2Mjc0NDQwMDgsInJvbGUiOiIxIiwiZXhwIjoxNjI5NjU3OTMwfQ.4P3tY5xUkKdUCHSlXliRnDDdrHK_cQNtskZgl6kTbWY",
		Width:    "100",
		Height:   "200",
		Duration: 11,
	}

	b1, e := json.Marshal(fi)
	fmt.Println(e)
	fmt.Println("这是 json 数据", string(b1))

	resp := ss.AddFile(string(b1))
	fmt.Println("这是返回的数据 =", resp)

}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
