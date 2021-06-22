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
	var art vo.ArticleCancelLikeParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	//查询是否存在记录
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(art.Token)
	if !b {
		return errors.New("token 失效")
	}
	//userid:=claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	//查询数据
	stmt, err := db.DB.Prepare("UPDATE article_like set is_like=? where id=?")
	if err != nil {
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return errors.New("更新数据失败")
	}
	res, err := stmt.Exec(int64(0), art.Id)
	if err != nil {
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}
	affect, err := res.RowsAffected()
	if affect == 0 {
		sugar.Log.Error("update article_like is failed.Err is ", err)
		return err
	}
	return nil
}
