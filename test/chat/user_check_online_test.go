package chat

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/test/myipfs"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestCheckUserOnline(t *testing.T) {
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

	req := vo.FriendCheckOnlineParams{
		Token:     token,
		FriendIds: []string{"416203556291354624"},
	}
	value, _ := json.Marshal(req)

	ss := Testdb(d)

	node, err := myipfs.GetIpfsNode("/Users/apple/projects/clouddb/test/chat/.ipfs")
	if err != nil {
		sugar.Log.Info("xxxxx----", err)
		panic(err)
	}

	// // h2ID, _ := peer.Decode("12D3KooWS8qWyGimuUgDjakUFGJkDgvGYcMEjnj5xqojeDwf1rZm")
	// h2ID, _ := peer.Decode("12D3KooWMUCCUigkLYryEJpGC1DdnJV87x8GozccreW2SVgK7KXW")

	// addr, err := node.DHT.FindPeer(context.Background(), h2ID)
	// if err != nil {
	// 	fmt.Println("find peer err:", err)
	// }

	// fmt.Println("-------addr:", addr)

	time.Sleep(20 * time.Second)

	resp := ss.FriendCheckOnline(node, string(value))
	t.Log("获取返回的数据 :=  ", resp)

	select {}
}
