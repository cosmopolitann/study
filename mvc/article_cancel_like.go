package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//朋友圈点赞
func ArticleCancelLike(db *Sql, value string) error {
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
	sugar.Log.Info("~~~~~  ArticleCancelLike   Method     End ~~~~~~")
	return nil
}
