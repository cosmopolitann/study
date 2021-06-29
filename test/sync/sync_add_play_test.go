package sync

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
)

func TestSyncAddPlay(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "../../tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	log.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)
	//	value:=`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","content":"三国"}
	//`
	// syncValue := `{"method":"SyncUser","data":{"id":"4324","peerId":"124","name":"20","phone":"1889","sex":"1","nickName":"nick"}}`
	syncValue := `{"id":"我想睡觉10"}`
	e = ss.SyncArticlePlay(syncValue)
	if e != nil {
		fmt.Println("e:=", e)
	}
	log.Println("方法不匹配")
}
