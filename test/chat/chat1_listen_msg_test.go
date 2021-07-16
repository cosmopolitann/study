package chat

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/myipfs"
	"github.com/cosmopolitann/clouddb/sugar"
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

	token, _ := jwt.GenerateToken("414207114215428096", "peerid", "name", "phone", "nickname", "img", 0, 1, 1, 30*24*60*60)

	fmt.Println(token)
	sugar.Log.Info("token: ", token)

	var cl ChatLister

	node, err := myipfs.GetIpfsNode("/Users/apple/projects/clouddb/test/chat/.ipfs")
	if err != nil {
		sugar.Log.Info("xxxxx----", err)
		panic(err)
	}

	err = ss.ChatListenMsgBlocked(node, token, &cl)
	fmt.Println(err)

	select {}

}

type ChatLister struct{}

func (cl *ChatLister) HandlerChat(abc string) {
	fmt.Println("1111", abc, "2222")
}
