package chat

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatRecordList(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Projects/clouddb/tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	fmt.Println(" Ping is failed,err:=", e)
	ss := Testdb(d)

	token, _ := jwt.GenerateToken("414202692580151296", "peerid", "name", "phone", "nickname", "img", 0, 1, 1, 30*24*60*60)

	req := vo.ChatRecordListParams{
		FromId:     "414202692580151296",
		Token:      token,
		CustomerId: "414202692580151333",
	}

	value, _ := json.Marshal(req)

	resp := ss.ChatRecordList(string(value))
	fmt.Println("获取返回的数据 :=  ", resp)

}
