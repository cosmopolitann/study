package sql

import (
	"io/ioutil"

	"github.com/cosmopolitann/clouddb/sugar"

	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestImportArticlTable(t *testing.T) {
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

	//stmt.QueryRow()
	out,
	ioutil.WriteFile("DB.sql", out, 0644)
	if err != nil {
		fmt.Println("err:=", err)

	}
	fmt.Println(res)
}
