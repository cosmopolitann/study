package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"strconv"
	"time"
)

func AddArticle(ipfsNode *ipfsCore.IpfsNode,db *Sql, value string) error {
	sugar.Log.Info(" ----  AddArticle Method ----")
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return errors.New(" Marshal article params is failed. ")
	}
	sugar.Log.Info("Marshal article params data : ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.Err: ", err)
		return errors.New(" Insert into article table is failed. ")
	}
	sid := strconv.FormatInt(id, 10)
	//stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error(" Insert into article  is Failed.", err)
		return errors.New(" Execute query article table is failed. ")
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New(" Insert into article table is failed. ")
	}

	// publish msg
	var ok bool
	topic:="/db-online-sync"
	var tp *pubsub.Topic
	ctx := context.Background()
	if tp,ok = Topicmp["/db-online-sync"];ok == false {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			return err
		}
		Topicmp[topic] = tp
	}
	sugar.Log.Info("Publish topic name :","/db-online-sync")
    //step 1
 	//query a article data
	var dl vo.ArticleResp
	rows, err := db.DB.Query("SELECT * from article where id=?;",sid)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return errors.New(" Sync query article table is failed. ")
	}
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.ShareNum, &dl.PlayNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}
	}
	var g PubSyncArticle
	g.Data = dl
	g.Type = "receiveArticleAdd"
	g.FromId=ipfsNode.Identity.String()
   //struct => json
	jsonBytes, err := json.Marshal(g)
	if err != nil {
		sugar.Log.Error("Marshal struct => json is failed.")
	return err
	}
	sugar.Log.Info("Forward the data to the public gateway.data:=",string(jsonBytes))
	err = tp.Publish(ctx,jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish info failed.Err:", err)
		return err
	}
	sugar.Log.Info("---  Publish to other device  ---")
	sugar.Log.Info(" ----  AddArticle Method  End ----")
	return nil
}
type PubSyncArticle struct {
	Type string `json:"type"`
	Data vo.ArticleResp `json:"data"`
	FromId string `json:"from"`
}
//

func ArticleList(db *Sql, value string) ([]Article, error) {
	sugar.Log.Info(" ----  ArticleList Method   ----")
	var art []Article
	var result vo.ArticleListParams
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return art,err
	}
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return art, errors.New(" Token is invalid. ")
	}
	userid := claim["UserId"]
	r := (result.PageNum - 1) * result.PageSize
	sugar.Log.Info("r:=", r)
	sugar.Log.Info("Claim:=", claim)
	sugar.Log.Info("userid :=", userid)
	sugar.Log.Info("Marshal data: ", result)
	sugar.Log.Info("PageNum:= ", result.PageNum)
	sugar.Log.Info("PageSize:= ", result.PageSize)
	//这里 要修改   加上 where  参数 判断
	rows, err := db.DB.Query("SELECT * FROM article where user_id=? order by ptime Desc limit ?,?",userid,r,result.PageSize)
	if err != nil {
		sugar.Log.Error("Query article table is failed.Err:", err)
		return art, errors.New(" Query article list is failed.")
	}
	for rows.Next() {
		var dl Article
		var userId interface{}
		var k=""
		err = rows.Scan(&dl.Id, &userId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		if userId==nil{
			dl.UserId=k
		}else {
			dl.UserId=userId.(string)
		}
		sugar.Log.Info("Query a data from article once.", dl)
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Query  article  is Failed.", err)
		return art, err
	}
	sugar.Log.Info("Query  article list is successful.")
	sugar.Log.Info(" ----  ArticleList  Method  End ----")
	return art, nil

}


func ArticleAddTest(db *Sql, value string) error{
	sugar.Log.Info(" ----  AddArticle Method ----")
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return errors.New(" Marshal article params is failed. ")
	}
	sugar.Log.Info("Marshal article params data : ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.Err: ", err)
		return errors.New(" Insert into article table is failed. ")
	}
	sid := strconv.FormatInt(id, 10)
	//stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error(" Insert into article  is Failed.", err)
		return errors.New(" Execute query article table is failed. ")
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New(" Insert into article table is failed. ")
	}
	return  nil
}