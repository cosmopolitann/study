package chat

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/myipfs"
	"github.com/cosmopolitann/clouddb/sugar"

	"github.com/cosmopolitann/clouddb/jwt"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatListenMsgBlock(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/projects/clouddb/tables/xiaolong.db")
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

	token, _ := jwt.GenerateToken("416203556291354624", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	var cl ChatListerBlocked

	node, err := myipfs.GetIpfsNode("/Users/apple/projects/clouddb/test/chat/.ipfs")
	if err != nil {
		sugar.Log.Info("xxxxx----", err)
		panic(err)
	}

	fmt.Println("---\n", node.PeerHost.ID().Pretty(), "\n----")

	// go func() {
	// 	sli := []string{
	// 		"416203556291354624",
	// 		"416418922095452160",
	// 		"416203557629337600",
	// 	}

	// 	for _, t := range sli {
	// 		time.Sleep(20 * time.Second)
	// 		token = ""
	// 		if t != "" {
	// 			token, _ = jwt.GenerateToken(t, "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)
	// 		}
	// 		ss.ChatListenMsgUpdateUser(node, token)
	// 	}

	// }()

	err = ss.ChatListenMsgBlocked(node, token, &cl)
	if err != nil {
		t.Error(err)
	}

}

type ChatListerBlocked struct{}

func (cl *ChatListerBlocked) HandlerChat(abc string) {
	fmt.Println("handlerChat----\n", abc, "2222-----")
}
