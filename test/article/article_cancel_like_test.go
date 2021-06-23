package article

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
	"testing"
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

	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","id":"414167149750259712"}
`
	t.Log("request value :=", value)
	resp := ss.ArticleCancelLike(value)
	t.Log("result:=", resp)
}
func Testdb1(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
