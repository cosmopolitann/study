package article

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/test/myipfs"
	_ "github.com/mattn/go-sqlite3"
)

func TestArticleCancelLike(t *testing.T) {
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
	sugar.Log.Info(" Ping is failed,err:= ", e)
	ss := Testdb1(d)
	// request json  params
	// test 1

	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTY5ODQ1NDUwNjIwMzEzNjAiLCJleHAiOjE2MjYzNTUxMTl9.Ko9C6ojPzShQ3BSP_ASa602EUjD27trRO_11zaV4hCY","id":"张柏芝30"}
`
	ipfsNode, err := myipfs.GetIpfsNode("/Users/apple/winter/D-cloud/test/ipfs")
	if err != nil {
		fmt.Println("err:=", err)
	}
	t.Log("request value :=", value)

	resp := ss.ArticleCancelLike(ipfsNode, value)

	t.Log("result:=", resp)
}
func Testdb1(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
