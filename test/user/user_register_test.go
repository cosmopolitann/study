package user

import (
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"

	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestUserRegister(t *testing.T) {

	//{"id":"4324","peerId":"124","name":"20","phone":1,"sex":"1","nickName":"nick"}

	// value := `{"id":"43243421","peerId":"Q1w213e1233221","name":"20","phone":"12233456","sex":"1","nickName":"nick","img":"123"}`
	//resp:= ss.UserAdd(string(b1)
	// resp := ss.UserRegister(value)
	// fmt.Println("这是返回的数据 =", resp)
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
	value := `{"id":"43243421","peerId":"QM00000111111","name":"20","phone":"009900099","sex":1,"nickName":"nick","img":"123"}`
	resp := ss.AddUserTest(string(value))
	fmt.Print("resp:===", resp)
}
func Testdb(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}

func TestExportUser(t *testing.T) {

}
