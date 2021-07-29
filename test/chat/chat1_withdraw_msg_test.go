package chat

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatWithdrawMsg(t *testing.T) {
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
	err = d.Ping()
	if err != nil {
		panic(err)
	}

	token, _ := jwt.GenerateToken("411647506288480256", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	req := vo.ChatWithdrawMsgParams{
		MsgId:  "412578662399873024",
		FromId: "411647506288480256",
		ToId:   "411642059200401408",
		Token:  token,
	}

	value, _ := json.Marshal(req)

	ss := Testdb(d)

	resp := ss.ChatWithdrawMsg(nil, string(value))
	t.Log("获取返回的数据 :=  ", resp)

}
