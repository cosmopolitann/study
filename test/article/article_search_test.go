package article

import (
	"database/sql"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

func TestArticleSearch(t *testing.T) {
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
	sugar.Log.Info(" Ping is failed,err:= ", e)
	ss := Testdb(d)
	// request json  params
	// test 1
	value := `{"pageSize":10,"pageNum":1,"title":"我"}
`
	t.Log("request value :=", value)


	resp1 := ss.ArticleSearch(value)

	t.Log("result1:=", resp1)

	//	// test 2
	//	value2:=`{"id":"411285804581654528"}
	//`
	//	t.Log("request value :=",value2)
	//	resp2:= ss.ArticlePlayAdd(value2)
	//	t.Log("result:=",resp2)

}
