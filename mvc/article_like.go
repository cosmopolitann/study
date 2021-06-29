package mvc

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
)

//朋友圈点赞

func AddArticleLike(db *Sql, value string) error {
	sugar.Log.Info("~~~~  AddArticleLike   Method  ~~~~~")
	var dl ArticleLike
	var art vo.ArticleGiveLikeParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//
	//check token is valid.
	claim, b := jwt.JwtVeriyToken(art.Token)
	if !b {
		return errors.New("token 失效")
	}
	userid := claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	//query data.
	rows, err := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(article_id,'null'),IFNULL(is_like,0) FROM article_like where article_id=? and user_id=?", art.Id, userid)
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

	if dl.Id == "" {
		//insert a new entire.
		id := utils.SnowId()
		stmt, err := db.DB.Prepare("INSERT INTO article_like (id,user_id,article_id,is_like) values(?,?,?,?)")
		if err != nil {
			sugar.Log.Error("Insert into article table is failed.", err)
			return err
		}
		sid := strconv.FormatInt(id, 10)
		stmt.QueryRow()
		res, err := stmt.Exec(sid, userid, art.Id, int64(1))
		if err != nil {
			sugar.Log.Error("Insert into article_like  is Failed.", err)
			return err
		}
		sugar.Log.Info("Insert into article_like  is successful.")
		l, _ := res.RowsAffected()
		if l == 0 {
			return errors.New(" Insert data into article_like table is failed. ")
		}
		sugar.Log.Info("~~~~  AddArticleLike   Method   End ~~~~~")
		return nil
	} else {
		//更新字段  is_ike = 1
		stmt, err := db.DB.Prepare("update article_like set is_like=? where article_id=? and user_id=?")
		if err != nil {
			sugar.Log.Error("update  article_like  is Failed.", err)
			return err
		}
		res, err := stmt.Exec(int64(1), art.Id, userid)
		if err != nil {
			sugar.Log.Error("update  article_like  is Failed.", err)
			return err
		}
		affect, err := res.RowsAffected()
		if affect == 0 {
			sugar.Log.Error("update article_like  is Failed.", err)
			return err
		}
		sugar.Log.Info("~~~~  AddArticleLike   Method   End ~~~~~")
		return nil
	}

}
