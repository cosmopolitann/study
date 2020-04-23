# 连接Mongo

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

