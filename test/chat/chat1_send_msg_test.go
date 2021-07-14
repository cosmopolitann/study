package chat

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	_ "github.com/mattn/go-sqlite3"
)

func TestChatSendMsg(t *testing.T) {
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

	token, _ := jwt.GenerateToken("414443377656860672", 30*24*60*60)

	req := vo.ChatSendMsgParams{
		RecordId:    "414443377656860672_414455684399108096",
		ContentType: 2,
		Content:     "content 22222222",
		FromId:      "414443377656860672",
		ToId:        "414455684399108096",
		Token:       token,
	}
	value, _ := json.Marshal(req)

	ss := Testdb(d)

	resp := ss.ChatSendMsg(nil, string(value))
	t.Log("获取返回的数据 :=  ", resp)

}
