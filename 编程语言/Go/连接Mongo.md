# 连接Mongo



## 1.连接的 mgo 包

```go
package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
)

const URL = "127.0.0.1:27017" //mongodb的地址

func main() {
	session, err := mgo.Dial(URL) //连接服务器
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")
	fmt.Println(session)

	c := session.DB("test").C("people") //选择ChatRoom库的account表

	jsonStr := `{
    "id": 1,
    "jsonrpc": "2.0",
    "result": {
        "author": "0xe79fead329b69540142fabd881099c04424cc49f",
        "extraData": "0xd68094302e392e312b2b35346466524c696e7578474e55",
        "gasLimit": "0x100000000",
        "gasUsed": "0x6c1e",
        "hash": "0x0b5c86108b16745971748a16b00eab67a043361916c821944328920a71b6a0f9",
        "logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
        "miner": "0xe79fead329b69540142fabd881099c04424cc49f",
        "number": "0x1d",
        "parentHash": "0x1a0e6c1953897d4562b94c0cef7397bb828ad279bca40a55c01042543713dc86",
        "receiptsRoot": "0xbde664cb5858c9cebab72d6f8d91dcddf8f40c1f88a1eb5f302ebe4a5f22331e",
        "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
        "sign": "0xcd1b66777cd96b58858026c0f203d840ace9d288b9e246ea8bd707ef8647808f1446c3019e0665486f557209669e0ac19c35cb3c96621a948561ed19ff296d1d01",
        "size": "0x0",
        "stateRoot": "0x35bd57206e933c6fdfc0e24fea42e8819286924f1572cd625c8c96e53f7e7742",
        "timestamp": "0x17196cd8e90",
        "totalDifficulty": "0x258",
        "transactions": [
            {
                "baseGas": "0x53d8",
                "blockHash": "0x0b5c86108b16745971748a16b00eab67a043361916c821944328920a71b6a0f9",
                "blockNumber": "0x1d",
                "from": "0xe79fead329b69540142fabd881099c04424cc49f",
                "gas": "0x872b580",
                "gasPrice": "0x5",
                "hash": "0xf5e40bf58d2d578231a8344db2baec3aea739c50154d384b25b11b4fbf7415d4",
                "input": "0x1e18dc450000000000000000000000000000000000000000000000000000000000000001",
                "nonce": "0xa",
                "r": "0x464e04cbb0bcd9dc3ad73ff2ee8a0dc12ee073f0f11b89c71e444758f5c241bd",
                "s": "0x67f38343de24a647d5673486805546330e71f9a0be2e930d77aa28a14c8f4137",
                "to": "0x8ad0955c90382d6f9948f331841be47abd938af3",
                "transactionIndex": "0x0",
                "transactionRlp": "0xf8850a05840872b580948ad0955c90382d6f9948f331841be47abd938af380a41e18dc45000000000000000000000000000000000000000000000000000000000000000126a0464e04cbb0bcd9dc3ad73ff2ee8a0dc12ee073f0f11b89c71e444758f5c241bda067f38343de24a647d5673486805546330e71f9a0be2e930d77aa28a14c8f4137",
                "v": "0x1",
                "value": "0x0"
            }
        ],
        "transactionsRoot": "0x705516727e60591d3d6cba34779ad877f2425207bb1091eda1f841a3816e4f2f",
        "uncles": []
    }
}
        `
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)

	c.Insert(mapResult)
	//
	//c.Insert(map[string]interface{}{"id": 7, "name": "tongjh", "age": 25}) //增
	//
	//objid := bson.ObjectIdHex("55b97a2e16bc6197ad9cad59")
	//
	//c.RemoveId(objid) //删除
	//
	//c.UpdateId(objid, map[string]interface{}{"id": 8, "name": "aaaaa", "age": 30}) //改
	//
	//var one map[string]interface{}
	//c.FindId(objid).One(&one) //查询符合条件的一行数据
	//fmt.Println(one)
	//
	//var result []map[string]interface{}
	//c.Find(nil).All(&result) //查询全部
	//fmt.Println(result)
}

```





https://blog.csdn.net/tamtian/article/details/105786611

//mongo



```go
package model

import (
	"brc_scan_block/log"
	"gopkg.in/ini.v1"
	"gopkg.in/mgo.v2"
)

var Session *mgo.Session

var C1 *mgo.Database

func InitDB(cfg *ini.File) {

	ip := cfg.Section("Mongo").Key("ip").String()
	port := cfg.Section("Mongo").Key("port").String()
	URL := ip + ":" + port
	log.Info.Println("URL=", URL)
	//mongodb://test:8888@127.0.0.1:28015/database?authSource=admin
	//[mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]

	//connect
	url := cfg.Section("Mongo").Key("url").String()
	log.Info.Println("URL=", url)

	var Session, err = mgo.Dial(URL)
	if err != nil {
		panic(err)
	}
	log.Info.Println("Connection successful！")

	//use database test
	log.Info.Println("Database is test.")
	C1 = Session.DB("test")

}

```



## 2.连接的 官方drive 包

```go
package model

import "C"
import (
	"brc_scan_block/log"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
)

var Client *mongo.Client

var C *mongo.Database

func InitDB_Mgo(cfg *ini.File) {
	// Read configuration.
	// url := cfg.Section("Mongo").Key("url").String()
	//log.Info.Println("连接Mongo的url===", url)

	//localhost
	url := "mongodb://localhost:27017"
	// Set client connection configuration.
	clientOptions := options.Client().ApplyURI(url)

	// connect MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error.Println(err)
	}

	// ping
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error.Println(err)
	}
	log.Info.Println("Connection mongo is successfully!")

	//use databases
	C = client.Database("test")

	log.Info.Println("Select database is test.")

}

func InitDB_Mgo1() {

	// Set client connection configuration.
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connect MongoDB
	var Client1, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Error.Println(err)
	}

	// ping
	err = Client1.Ping(context.TODO(), nil)
	if err != nil {
		log.Error.Println(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection := Client1.Database("test").Collection("m2")
	insertResult, err := collection.InsertOne(context.TODO(), map[string]interface{}{
		"mgo2": "drive2",
	})

	fmt.Println("Inserted a single document: ", insertResult)
	if err != nil {
		log.Error.Println(err)
	}

}

```

