package cloud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestAddFile(t *testing.T) {
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
	var fi = vo.CloudAddFileParams{
		Id:       "411580511585046528",
		FileName: "我爱成都1",
		ParentId: "0",
		FileCid:  "Qm123",
		FileSize: 100,
		FileType: 0,
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTY5ODQ1NDUwNjIwMzEzNjAiLCJleHAiOjE2MjYzNTUxMTl9.Ko9C6ojPzShQ3BSP_ASa602EUjD27trRO_11zaV4hCY",
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
