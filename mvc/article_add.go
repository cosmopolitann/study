package mvc

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func AddArticle(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string, path string) error {
	sugar.Log.Info(" ----  AddArticle Method ----")
	sugar.Log.Info(" ----  Path :", path)
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return errors.New(" Marshal article params is failed. ")
	}
	sugar.Log.Info("Marshal article params data : ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size) values (?,?,?,?,?,?,?,?,?,?,?,?,?)")
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

	//--------------- publish msg ----------------
	// var ok bool
	topic := "/db-online-sync"
	var tp *pubsub.Topic
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
	sugar.Log.Info("Publish topic name :", "/db-online-sync")
	//step 1
	//query a article data
	var dl vo.ArticleResp
	err = db.DB.QueryRow("SELECT id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article where id=?;", sid).Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
	if err != nil && err != sql.ErrNoRows {
		sugar.Log.Error("Query article failed.Err is", err)
		return err
	}
	//

	//query user info.
	var dl1 vo.RespSysUser
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') from sys_user where id=?", art.UserId)
	if err != nil {
		sugar.Log.Error("AddUser Query data is failed.Err is ", err)
		return err
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl1.Id, &dl1.PeerId, &dl1.Name, &dl1.Phone, &dl1.Sex, &dl1.Ptime, &dl1.Utime, &dl1.NickName, &dl1.Img)
		if err != nil {
			sugar.Log.Error("AddUser Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info(" AddUser Query a entire data is ", dl)
	}

	//the first step.
	var s3 UserAd
	s3.Type = "receiveUserRegister"
	s3.Data = dl1
	s3.FromId = ipfsNode.Identity.String()
	//marshal UserAd.
	//the second step
	sugar.Log.Info("--- second step ---")

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Error("Publish msg is failed.Err:", err)
		return err
	}
	sugar.Log.Info("Frwarding information:=", string(jsonBytes))
	sugar.Log.Info("Local PeerId :=", ipfsNode.Identity.String())
	//the  third  step .
	sugar.Log.Info("--- third step ---")

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
		return err
	}

	//----
	var g PubSyncArticle
	g.Data = dl
	g.Type = "receiveArticleAdd"
	g.FromId = ipfsNode.Identity.String()
	//struct => json
	jsonBytes, err = json.Marshal(g)
	if err != nil {
		sugar.Log.Error("Marshal struct => json is failed.")
		return err
	}
	sugar.Log.Info("Forward the data to the public gateway.data:=", string(jsonBytes))

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish info failed.Err:", err)
		return err
	}
	sugar.Log.Info("---  Publish to other device  ---")
	//

	sugar.Log.Info("~~~~  Publish msg is successful.   ~~~~  ")

	// 写入文件
	sugar.Log.Info("---  write sql to file.  ---")
	//
	f1, err1 := os.OpenFile(path+"update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
	if err1 != nil {
		sugar.Log.Error(" Create update is failed. Err :", err1)
	}

	sugar.Log.Info("----- Local file is exist.  ----")

	// 拼接字符串 sql 语句

	sql := fmt.Sprintf("INSERT INTO article (id,user_id,accesstory,accesstory_type,text,tag,ptime,play_num,share_num,title,thumbnail,file_name,file_size) values ('%s','%s','%s',%d,'%s','%s',%d,%d,%d,'%s','%s','%s','%s')\n", sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)

	_, err = f1.WriteString(sql)
	if err != nil {
		sugar.Log.Error("-----  Write sql to update file is failed.Err:  ----", err)
	}
	sugar.Log.Info("-----sql :-----", sql)

	sugar.Log.Info("-----  Write sql to file is successful~~ ----")
	sugar.Log.Info(" ----  AddArticle Method  End ----")
	return nil
}

type PubSyncArticle struct {
	Type   string         `json:"type"`
	Data   vo.ArticleResp `json:"data"`
	FromId string         `json:"from"`
}

//

func ArticleList(db *Sql, value string) ([]Article, error) {
	sugar.Log.Info(" ----  ArticleList Method   ----")
	var art []Article
	var result vo.ArticleListParams
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err:", err)
		return art, err
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
	rows, err := db.DB.Query("SELECT IFNULL(b.is_like,0),a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),(SELECT COUNT( * ) FROM article_like AS c WHERE c.article_id = b.article_id ) as sum FROM article as a LEFT JOIN article_like as b on a.id=b.article_id where a.user_id=? order by ptime Desc limit ?,?", userid, r, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query article table is failed.Err:", err)
		return art, errors.New(" Query article list is failed.")
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		var dl Article
		var userId interface{}
		var k = ""
		err = rows.Scan(&dl.IsLike, &dl.Id, &userId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.LikeNum)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		if userId == nil {
			dl.UserId = k
		} else {
			dl.UserId = userId.(string)
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

func ArticleAddTest(db *Sql, value string) error {
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
	return nil
}
