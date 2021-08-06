package chat

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cosmopolitann/clouddb/jwt"
)

func TestChatCreateRecord(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/Projects/clouddb/tables/xiaolong.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	err = d.Ping()
	if err != nil {
		panic(err)
	}

	token, _ := jwt.GenerateToken("411642059200401408", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	fmt.Println(token)

	req := vo.ChatAddRecordParams{
		Name:   "Record Name 2222",
		FromId: "411642059200401408",
		ToId:   "411647506288480256",
		Token:  token,
	}

	value, _ := json.Marshal(req)

	ss := Testdb(d)
	resp := ss.ChatCreateRecord(nil, string(value))
	t.Log("获取返回的数据 :=  ", resp)

}
