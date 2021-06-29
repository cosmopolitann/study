package user

import (
	"github.com/cosmopolitann/clouddb/sugar"

	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestUserLogin(t *testing.T) {
	//str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIzMjczNzM2NTI2NDI0NzUwMDgiLCJjdGltZSI6MTYyNDM1Mjc2MTEwMCwiaWF0IjoxNjI0MzUyNzYxfQ.zduJoHT-qM6bvySnXWVImrieKx-MdO3bYH4PSjL-5wo"
	//a,_ := jwt.JwtVeriyToken(str)
	//fmt.Println("claim:",a)
	//return
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

	//{"id":"4324","peerId":"124","name":"20","phone":1,"sex":"1","nickName":"nick"}
	value := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTYzMzgyMjM0MjY0NDEyMTYifQ.0XcXav8J2pvo1Tp3S2y8GhIM5oNgZA7982l7FtNt1Yc`
	//resp:= ss.UserAdd(string(b1)

	resp := ss.UserLogin(value)

	fmt.Println("这是返回的数据 =", resp)

}
