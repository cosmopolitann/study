package sync

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

func TestQueryAllData(t *testing.T) {
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
	value := `{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMjMiLCJwZWVySWQiOiJwZXJyIiwibmFtZSI6Im5hbWUiLCJwaG9uZSI6InBob25lIiwic2V4IjoxLCJuaWNrTmFtZSI6Im5pY2siLCJpbWciOiJpbWciLCJleHAiOjE2MjU3NTAxMzd9.I6J8fE1SbSiNiyd-WIiSawRFQ_tGA9PEt0jHNKyxVxo"}`
	path := "/Users/apple/winter/offline/querydata"

	resp := ss.SyncQueryAllData(value, path)
	fmt.Println("这是返回的数据 =", resp)

}
func TestDatabaseMigration(t *testing.T) {
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
	token := `{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI"}`
	resp := ss.SyncDatabaseMigration(token, "/Users/apple/winter/offline/", "QmaNTroLZEX77UwS5GixEp7jX12sK5mBjkTAdGagYTRy9k")
	fmt.Println("这是返回的数据 =", resp)

}
