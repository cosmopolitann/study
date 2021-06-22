package main

import (
	"database/sql"
	"testing"

	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
)

func TestOffLineSyncData(t *testing.T) {

	t.Run("offline sync", func(t *testing.T) {
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
		t.Log(" Ping is failed,err:=", e)

		ss := Testdb(d)
		resp := ss.OfflineSync("123")
		t.Log("result:=", resp)
	})

}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
