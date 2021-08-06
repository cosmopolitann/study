package chat

import (
	"database/sql"
	"log"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
)

func TestChatRecord(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/xiaolong.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	log.Println(" Ping is failed,err:=", e)
	// 	ss := Testdb(d)
	// 	value := `{"Id":"2342","RecordName":"sfwer","RecordTopic":"20","RecordImg":"123","CreateBy":"1","LastMsg":"sfs"}
	// `
	// 	resp := ss.AddChatRecord(value)
	// 	log.Println("这是返回的数据 =", resp)
}
