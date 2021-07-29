package sql

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
)

func TestImportArticlTable(t *testing.T) {

	sugar.InitLogger()

	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/clouddb/tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")

	result, err := d.Exec("alter table sys_user add column role varchar(64) not null default(1)")
	if err != nil {
		fmt.Println("alter is err,err:=", err)

	}
	fmt.Println("result:=", result)
	result2, err := d.Exec("alter table article add column external_href text")
	if err != nil {
		fmt.Println("alter is err,err:=", err)

	}
	fmt.Println("result:=", result2)

	result3, err := d.Exec("INSERT INTO sys_user (id, peer_id, name, phone, sex, ptime, utime, nickname, img, role) VALUES ('天天想你啊', '', '星河飞天-人工客服12323', '', 0, 1627444008, 1627444008, '人工客服1', '', '1')")
	if err != nil {
		fmt.Println("alter is err,err:=", err)

	}
	fmt.Println("result3:=", result3)

}
