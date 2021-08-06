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

func TestFriendUpdateNickname(t *testing.T) {
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

	token, _ := jwt.GenerateToken("416203557629337600", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	req := vo.FriendUpdateNicknameParams{
		Token:    token,
		FriendId: "416418921625690112",
		Nickname: "sssb122222222",
	}
	value, _ := json.Marshal(req)

	ss := Testdb(d)

	resp := ss.FriendUpdateNickname(string(value))
	t.Log("获取返回的数据 :=  ", resp)
}
