package mvc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
)

//朋友圈点赞
func ArticleCancelLike(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) error {
	sugar.Log.Info("~~~~~  ArticleCancelLike   Method   ~~~~~~")
	var art vo.ArticleCancelLikeParams
	//unmarshal params info.
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//check token is valid.
	claim, b := jwt.JwtVeriyToken(art.Token)
	if !b {
		return errors.New(" Token is invalid. ")
	}
	//userid:=claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	//First,query data from article_like table. where id=?,
	//then update it.
	stmt, err := db.DB.Prepare("UPDATE article_like set is_like=? where id=?")
	if err != nil {
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(0), art.Id)
	if err != nil {
		sugar.Log.Error("update exec article_like is failed.Err is ", err)
		return err
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}
	// query

	//publish msg
	topic := "/db-online-sync"
	sugar.Log.Info("发布主题:", "/db-online-sync")
	sugar.Log.Info("发布消息:", value)

	var dl ArticleLike

	ctx := context.Background()

	tp, ok := TopicJoin.Load(topic)
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}
		TopicJoin.Store(topic, tp)
	}
	rows, err := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(article_id,'null'),IFNULL(is_like,0) FROM article_like where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query article_like is failed.Err is ", err)
		return err
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.ArticleId, &dl.IsLike)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}
	}
	sugar.Log.Info("--- 开始 发布的消息 ---")

	sugar.Log.Info("发布的消息:", value)
	//
	//第一步
	var s3 ArticleCanCelLike
	s3.Type = "receiveArticleCancelLike"
	s3.Data = dl
	s3.FromId = ipfsNode.Identity.String()

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Info("--- 开始 发布的消息 ---")
		return err
	}
	sugar.Log.Info("--- 解析后的数据 返回给 转接服务器 ---", string(jsonBytes))

	//============================
	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")

	//==
	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("发布错误:", err)
		return err
	}
	sugar.Log.Error("---  发布的消息  完成  ---")
	//---
	sugar.Log.Info("~~~~~  ArticleCancelLike   Method     End ~~~~~~")
	return nil
}

type ArticleCanCelLike struct {
	Type   string      `json:"type"`
	Data   ArticleLike `json:"data"`
	FromId string      `json:"from"`
}
