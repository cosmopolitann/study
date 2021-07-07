package main

import (
	"database/sql"
	"fmt"

	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/profile"
)

type Cloud struct {
	d mvc.Sql
}

func tt() (*Cloud, error) {
	//日志运行
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	d := mvc.Newdb("/Users/apple/winter/D-cloud/tables/foo.db")
	e := d.Ping()
	if e != nil {
		sugar.Log.Info(" 这是 Ping 的 err", e)
		return &Cloud{d: d}, e
	}
	sugar.Log.Info("创建数据库 完成")
	return &Cloud{d: d}, nil
}

func main() {
	//test
	// d := mvc.NTestNode("")
	// err := d.Add()
	// sugar.Log.Info("创建数据库失败，错误:", err)
	//example-folder
	// example_folder.InItipfs()
	// time.Sleep(time.Hour)
	//test
	defer profile.Start(profile.MemProfile).Stop()
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

	// ss := Testdb(d)
	// path := "/Users/apple/winter/offline/"
	// ss.OfflineSync(path)

}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
