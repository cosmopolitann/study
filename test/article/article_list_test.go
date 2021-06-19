package article

import (
	"database/sql"

	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

func TestArticleList(t *testing.T) {
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
	ss := Testdb2(d)

	//	value:=`{"id":"4324","userId":"123","accesstory":"20","accesstoryType":1,"text":"1","tag":"1","playNum":3,"title":"成都23","shareNum":4,"thumbnail":"刘亦菲3"}
	//`

	value := `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","pageSize":30000,"pageNum":1}
`
	resp := ss.ArticleList(value)

	//resp := (value)
	fmt.Println("这是返回的数据 =", resp)

}
