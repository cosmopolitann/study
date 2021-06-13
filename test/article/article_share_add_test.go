package article

import (
	"github.com/cosmopolitann/clouddb/sugar"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestAddArticleShareAdd(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.mvc")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	sugar.Log.Info(" Ping is failed,err:= ",e)
	ss:=Testdb(d)
	// request json  params
	// test 1
	value:=`{"id":"408575816247545856"}
`
	t.Log("request value :=",value)
	resp:= ss.ArticleShareAdd(value)
	t.Log("result:=",resp)


	// test 2
	value2:=`{"id":"408583113132740608"}
`
	t.Log("request value :=",value2)
	resp2:= ss.ArticleShareAdd(value2)
	t.Log("result:=",resp2)


}