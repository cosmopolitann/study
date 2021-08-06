package chat

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/test/myipfs"
	"github.com/cosmopolitann/clouddb/vo"
	"github.com/libp2p/go-libp2p-core/peer"
	_ "github.com/mattn/go-sqlite3"
)

func TestChatSendMsg(t *testing.T) {
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

	token, _ := jwt.GenerateToken("416418922095452160", "peerid", "name", "phone", "nickname", "img", "2", 0, 1, 1, 30*24*60*60)

	req := vo.ChatSendMsgParams{
		RecordId:    "416203556291354624_416418922095452160",
		ContentType: 2,
		Content:     "content 222222223333",
		FromId:      "416418922095452160",
		ToId:        "416203556291354624",
		Token:       token,
	}
	value, _ := json.Marshal(req)

	ss := Testdb(d)

	node, err := myipfs.GetIpfsNode("/Users/apple/projects/clouddb/test/chat/.ipfs")
	if err != nil {
		sugar.Log.Info("xxxxx----", err)
		panic(err)
	}

	// h2ID, _ := peer.Decode("12D3KooWS8qWyGimuUgDjakUFGJkDgvGYcMEjnj5xqojeDwf1rZm")
	h2ID, _ := peer.Decode("12D3KooWMUCCUigkLYryEJpGC1DdnJV87x8GozccreW2SVgK7KXW")

	addr, err := node.DHT.FindPeer(context.Background(), h2ID)
	if err != nil {
		fmt.Println("find peer err:", err)
	}

	fmt.Println("addr:", addr)

	resp := ss.ChatSendMsg(node, string(value))
	t.Log("获取返回的数据 :=  ", resp)

	select {}
}
