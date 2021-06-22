package article

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

//文章获取详情
func TestArticleQuery(t *testing.T) {
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
	ss := Testdb(d)
	// request json  params
	// test 1
	value := `{"id":"413271162890489856"}
`
	t.Log("request value :=", value)
	resp := ss.ArticleQuery(value)
	t.Log("result:=", resp)

	//	// test 2
	//	value2:=`{"id":"411555061567590400"}
	//`
	//	t.Log("request value :=",value2)
	//	resp2:= ss.ArticleQuery(value2)
	//	t.Log("result:=",resp2)

}
