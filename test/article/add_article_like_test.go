package article

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/test/myipfs"
	_ "github.com/mattn/go-sqlite3"
)

func TestAddArticleLike(t *testing.T) {
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
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)

	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTY5ODQ1NDUwNjIwMzEzNjAiLCJleHAiOjE2MjYzNTUxMTl9.Ko9C6ojPzShQ3BSP_ASa602EUjD27trRO_11zaV4hCY","id":"张柏芝30"}
`
	ipfsNode, err := myipfs.GetIpfsNode("/Users/apple/winter/D-cloud/test/ipfs")
	if err != nil {
		fmt.Println("err:=", err)
	}
	resp := ss.ArticleGiveLike(ipfsNode, value)
	fmt.Println("这是返回的数据 =", resp)
}
