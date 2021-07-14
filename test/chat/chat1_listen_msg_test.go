package chat

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"

	"github.com/cosmopolitann/clouddb/jwt"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatListenMsg(t *testing.T) {
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

	ss := Testdb(d)

	token, _ := jwt.GenerateToken("409330202166956089", "peerid", "name", "phone", "nickname", "img", 0, 1, 1, 30*24*60*60)

	fmt.Println(token)
	sugar.Log.Info("token: ", token)

	var cl ChatLister

	// node, err := ipfs.GetIpfsNode("/Users/apple/.ipfs")
	// if err != nil {
	// 	sugar.Log.Info("xxxxx----", err)
	// 	panic(err)
	// }

	resp := ss.ChatListenMsg(nil, token, &cl)
	t.Log("获取返回的数据 :=  ", resp)

	time.Sleep(time.Hour)

}

type ChatLister struct{}

func (cl *ChatLister) HandlerChat(abc string) {
	fmt.Println("1111", abc, "2222")
}
