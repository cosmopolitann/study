# Mongo聚合函数

###### mgo 官网  https://godoc.org/gopkg.in/mgo.v2#pkg-subdirectories

###### 这是mongo的 连接数据库，增删改查。

```go

package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const URL = "127.0.0.1:27017" //mongodb的地址
type User1 struct {
	Name string
	Home string
}

func main() {
	session, err := mgo.Dial(URL) //连接服务器
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")
	c := session.DB("test") //选择ChatRoom库的user表
	//查找数据

	//这里 的问题  就是   这个 ——id 字段  一定 需要
	//query := c.C("user").Find(bson.M{"_id": bson.ObjectIdHex("5ea78b2a6f8474c1e4c5f6bc")})
	query := c.C("user").Find(bson.M{"name": "123"})

	fmt.Println("query=", query)
	var s map[string]interface{}
	err = query.One(&s)
	if err != nil {
		fmt.Println(err)
	}
	//取出相应数据
	fmt.Println("s=", s)
	fmt.Printf("%T\n", s)




	d := s["home"]
	f := s["money"]

	fmt.Println(d, f)
	//插入一条数据
	//err = c.C("user").Insert(map[string]interface{}{
	//	"name": "lily",
	//	"age":  "18",
	//	"hobby": map[int]int{
	//		1: 1,
	//	},
	//})
	//	fmt.Println("插入成功")
	//插入一条数据  json 数据
	//	jsonstr := `{"id":123,
	//"name":"123"}`
	//
	//	var mapResult map[string]interface{}
	//	err = json.Unmarshal([]byte(jsonstr), &mapResult)
	//	if err != nil {
	//		fmt.Println("JsonToMapDemo err: ", err)
	//	}
	//	err = c.C("user").Insert(mapResult)
	//	fmt.Println("插入成功")

	//ObjectIdHex("5ea78b2a6f8474c1e4c5f6bc")

	//更新数据
	//err = c.C("user").Update(bson.M{"_id": bson.ObjectIdHex("5ea78b2a6f8474c1e4c5f6bc")}, bson.M{"$set": bson.M{"money": "10000"}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("更新成功")

	// 删除数据
	//err = c.C("user").Remove(bson.M{"_id": bson.ObjectIdHex("5ea79a656f8474c1e4c5fd30")})
	//fmt.Println("删除成功")

}

```

### 聚合函数

###### 1.查找 满足 name 是 123 的 文档 然后 数据都返回

```go

> db.test_data.aggregate([{$match:{name:"123"}}])

{ "_id" : ObjectId("5ea7c66b6f8474c1e4c602d6"), "name" : "123", "age" : "18", "address" : { "constrct" : "0x00000", "phone" : "0000000" } }
{ "_id" : ObjectId("5ea7c68d6f8474c1e4c60335"), "name" : "123", "age" : "18", "address" : { "constrct" : "0x1111111", "phone" : "11111111" } }
{ "_id" : ObjectId("5ea7c6996f8474c1e4c6033c"), "name" : "123", "age" : "18", "address" : { "constrct" : "0x222222222", "phone" : "22222222" } }
{ "_id" : ObjectId("5ea7c6a56f8474c1e4c60342"), "name" : "123", "age" : "18", "address" : { "phone" : "333333333", "constrct" : "0x33333333" } }
```

###### 2.查找 name 是  123 的 文档  然后  返回指定数据  比如  只返回 address   只返回 constrct。

###### constrct  是一个 数组集合。

```go
> db.test_data.aggregate([{$match:{name:"123"}},{$project:{"_id":0,"address":1}}])
{ "address" : { "constrct" : "0x00000", "phone" : "0000000" } }
{ "address" : { "constrct" : "0x1111111", "phone" : "11111111" } }
{ "address" : { "constrct" : "0x222222222", "phone" : "22222222" } }
{ "address" : { "phone" : "333333333", "constrct" : "0x33333333" } }
> db.test_data.aggregate([{$match:{name:"123"}},{$project:{"_id":0,"address":{"constrct":1}}}])
{ "address" : { "constrct" : "0x00000" } }
{ "address" : { "constrct" : "0x1111111" } }
{ "address" : { "constrct" : "0x222222222" } }
{ "address" : { "constrct" : "0x33333333" } }
>
```

