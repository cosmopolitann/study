package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func ArticlePlayAdd(ipfsNode *ipfsCore.IpfsNode,db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set play_num=? where id=?")
	if err != nil {
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.PlayNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	// publish msg
	topic:="/db-online-sync"
	sugar.Log.Info("发布主题:","/db-online-sync")
	sugar.Log.Info("发布消息:",value)
	//判断是否弃用
	var tp *pubsub.Topic
	var ok bool
	ctx := context.Background()
	if tp,ok = Topicmp["/db-online-sync"];ok == false {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			return err
		}
		Topicmp[topic] = tp
	}

	rows1, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}
	for rows1.Next() {
		err = rows1.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" update is failed .")
	}
	sugar.Log.Info("--- 开始 发布的消息 ---")

	sugar.Log.Info("发布的消息:", value)
	//
	//第一步
	var s3 PlayAdd
	s3.Type = "receiveArticlePlayAdd"
	s3.Data = dl
	s3.FromId=ipfsNode.Identity.String()

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Info("--- 开始 发布的消息 ---")
		return err
	}
	sugar.Log.Info("--- 解析后的数据 返回给 转接服务器 ---",string(jsonBytes))

	//============================
	err = tp.Publish(ctx,jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")

	//==
	err = tp.Publish(ctx,jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")

	return nil
}

// share add

type PlayAdd struct {
	Type string   `json:"type"`
	Data  Article `json:"data"`
	FromId string `json:"from"`
}


func ArticleShareAdd(ipfsNode *ipfsCore.IpfsNode,db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}


	//=====
	// publish msg
	topic:="/db-online-sync"
	sugar.Log.Info("发布主题:","/db-online-sync")
	sugar.Log.Info("发布消息:",value)
	//判断是否弃用
	var tp *pubsub.Topic
	var ok bool
	ctx := context.Background()
	if tp,ok = Topicmp["/db-online-sync"];ok == false {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			return err
		}
		Topicmp[topic] = tp

	}
	sugar.Log.Info("--- 开始 发布的消息 ---")

	sugar.Log.Info("发布的消息:", value)
	//
	//查询数据
	//查询数据

	//select the data is exist.
	rows1, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}

	for rows1.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" update is failed .")
	}
	//



	//
	//第一步
	var s3 ShareAdd
	s3.Type = "receiveArticleShareAdd"
	s3.Data = dl
	s3.FromId=ipfsNode.Identity.String()
	//

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Info("--- 开始 发布的消息 ---")
		return err
	}
	sugar.Log.Info("--- 解析后的数据 返回给 转接服务器 ---",string(jsonBytes))


	err = tp.Publish(ctx,jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")
	return nil
}

type ShareAdd struct {
	Type string `json:"type"`
	Data  Article `json:"data"`
	FromId string  `json:"from"`
}