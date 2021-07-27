package article

import (
	"database/sql"
	"encoding/json"

	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func TestArticleListUser(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/projects/clouddb/tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb2(d)

	req := vo.ArticleListUserParams{
		PageSize: 10,
		PageNum:  0,
		UserId:   "327719543623991296",
	}

	reqBytes, _ := json.Marshal(req)

	resp := ss.ArticleListUser(string(reqBytes))

	fmt.Println("这是返回的数据 =", resp)

}
