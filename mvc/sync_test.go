package mvc

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/myipfs"
	"github.com/cosmopolitann/clouddb/sugar"
	shell "github.com/ipfs/go-ipfs-api"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	_ "github.com/mattn/go-sqlite3"
)

//Test Post upload file.
func TestPostFormDataPublicgatewayFile(t *testing.T) {
	sugar.InitLogger()
	hash, err := PostFormDataPublicgatewayFile("/Users/apple/winter/D-cloud/sugar/", "remote")
	fmt.Println("hash=", hash)
	fmt.Println("err=", err)

}

func TestResolverIpnsAddress(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "", want: "", wantErr: true},
	}
	for _, tt := range tests {
		sugar.InitLogger()

		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolverIpnsAddress()

			t.Log("got:=", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolverIpnsAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolverIpnsAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDiffBetween(t *testing.T) {
	sugar.InitLogger()
	v1 := ""
	v2 := "QmeFiRN6Vreu2EvECv25yw75GqEdsGw1QF45w4bqGcHYF5_QmeVsha4XkdayJjpzD6ykJR6iXmF3VytyxE18sYCSCkgBR"
	d := FindDiffBetween(v1, v2)
	fmt.Println("d=", d)
}

// 创建 local 文件
func TestLocalNonexistent(t *testing.T) {

	sugar.InitLogger()

	err := LocalNonexistent("/Users/apple/winter/offline/")
	if err != nil {
		fmt.Println("err:=", err)
	}

}

func TestReadRemoteAndLocal(t *testing.T) {
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	// d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	path := "/Users/apple/winter/offline/"
	hash := "QmeFiRN6Vreu2EvECv25yw75GqEdsGw1QF45w4bqGcHYF5"

	gotV1, gotV2, err := ReadRemoteAndLocal(path, hash)
	if err != nil {
		fmt.Println("err:=", err)
	}
	fmt.Println("v1:=", gotV1)
	fmt.Println("v1:=", gotV2)
}

func TestLoopGetCidAndExcuteSql(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Info(" ~~~~~~~~~ Capture the panic ~~~~~~~~~~~~Err: ", err)
		} else {
			sugar.Log.Info("~~~~~~~~~~~~~~~   Normal ~~~~~~~~~~~~")
		}
	}()
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}
	db := Testdb(d)
	v1 := ""
	v2 := "QmeFiRN6Vreu2EvECv25yw75GqEdsGw1QF45w4bqGcHYF5_QmeVsha4XkdayJjpzD6ykJR6iXmF3VytyxE18sYCSCkgBR"
	diff := FindDiffBetween(v1, v2)
	path := "/Users/apple/winter/offline/"
	fmt.Println("diff:=", diff)
	fmt.Println("db:=", db)
	err = LoopGetCidAndExcuteSql(diff, path, &db)
	if err != nil {
		fmt.Println("err:=", err)
	}
	fmt.Println("=========")

}
func Testdb(sq *sql.DB) Sql {
	return Sql{DB: sq}
}

func TestUploadFile(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Info(" ~~~~~~~~~ Capture the panic ~~~~~~~~~~~~Err: ", err)
		} else {
			sugar.Log.Info("~~~~~~~~~~~~~~~   Normal ~~~~~~~~~~~~")
		}
	}()
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	hash := "QmTD5GfbjzznxAxKgn6H4sEUjHsvtXZd39TRgGY7nnVEo3"
	path := "/Users/apple/winter/offline/"
	//--- ipfs Node
	ipfsNode, err := myipfs.GetIpfsNode("/Users/apple/winter/D-cloud/test/ipfs")
	if err != nil {
		fmt.Println("err:=", err)
	}
	sugar.Log.Info("ipfsNode :", ipfsNode)

	// //publish  hash => public gateway.
	topic := "doudou"
	var tp *pubsub.Topic
	ctx := context.Background()
	tp, ok := TopicJoin.Load(topic)
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
		}
		TopicJoin.Store(topic, tp)
	}
	fmt.Println("topic :=", topic)
	fmt.Println("tp :=", tp)

	sugar.Log.Info("Publish topic name :", "doudou")
	got, err := UploadFile(path, hash)
	fmt.Println(err)
	fmt.Println(got)

	err = tp.Publish(ctx, []byte(got))
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
	}
	sugar.Log.Info("Publish is successful ~~~~~~")

	// recieve     msg

	// sub, err := tp.Subscribe()
	// if err != nil {
	// 	sugar.Log.Error("subscribe failed.", err)
	// }

	// for {
	// 	sugar.Log.Info("-----  Start Subscribe ------")
	// 	data, err := sub.Next(ctx)
	// 	if err != nil {
	// 		sugar.Log.Error("subscribe failed.", err)
	// 		continue
	// 	}
	// 	sugar.Log.Info("-----收到的 消息 ------", data)

	// }
	// select {}
}

func TestOffLineSyncData(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Info(" ~~~~~~~~~ Capture the panic ~~~~~~~~~~~~Err: ", err)
		} else {
			sugar.Log.Info("~~~~~~~~~~~~~~~   Normal ~~~~~~~~~~~~")
		}
	}()
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	hash := "QmTD5GfbjzznxAxKgn6H4sEUjHsvtXZd39TRgGY7nnVEo3"
	path := "/Users/apple/winter/offline/"
	//--- ipfs Node
	ipfsNode, err := myipfs.GetIpfsNode("/Users/apple/winter/D-cloud/test/ipfs")
	if err != nil {
		fmt.Println("err:=", err)
	}
	sugar.Log.Info("ipfsNode :", ipfsNode)

	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}
	db := Testdb(d)
	err = OffLineSyncData(&db, path, ipfsNode)
	if err != nil {
		sugar.Log.Error("  ~~~~~~    OffLineSyncData  ~~~~~   Err is", err)

	}

	// //publish  hash => public gateway.
	topic := "doudou"
	var tp *pubsub.Topic
	ctx := context.Background()
	tp, ok := TopicJoin.Load(topic)
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
		}
		TopicJoin.Store(topic, tp)
	}
	fmt.Println("topic :=", topic)
	fmt.Println("tp :=", tp)

	sugar.Log.Info("Publish topic name :", "doudou")
	got, err := UploadFile(path, hash)
	fmt.Println(err)
	fmt.Println(got)

	err = tp.Publish(ctx, []byte(got))
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
	}
	sugar.Log.Info("Publish is successful ~~~~~~")

}
