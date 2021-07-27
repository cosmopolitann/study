package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleList(db *Sql, value string) ([]ArticleAboutMeResp, error) {
	sugar.Log.Info(" ----  ArticleList Method   ----")
	var art []ArticleAboutMeResp
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
	userid := claim["id"]

	r := (result.PageNum - 1) * result.PageSize
	sugar.Log.Info("r:=", r)
	sugar.Log.Info("Claim:=", claim)
	sugar.Log.Info("userid :=", userid)
	sugar.Log.Info("Marshal data: ", result)
	sugar.Log.Info("PageNum:= ", result.PageNum)
	sugar.Log.Info("PageSize:= ", result.PageSize)
	//这里 要修改   加上 where  参数 判断
	rows, err := db.DB.Query("SELECT IFNULL(d.peer_id,''),IFNULL(d.name,''),IFNULL(d.phone,''),IFNULL(d.sex,0),IFNULL(d.nickname,''),IFNULL(d.img,''),IFNULL(b.is_like,0),a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),(SELECT COUNT( * ) FROM article_like AS c WHERE c.article_id = b.article_id ) as sum FROM article as a LEFT JOIN article_like as b on a.id=b.article_id LEFT JOIN sys_user as d on d.id=a.user_id where a.user_id=? order by a.ptime Desc limit ?,?", userid, r, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query article table is failed.Err:", err)
		return art, errors.New(" Query article list is failed")
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		var dl ArticleAboutMeResp
		var userId interface{}
		var k = ""
		err = rows.Scan(&dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName, &dl.Img, &dl.IsLike, &dl.Id, &userId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.LikeNum)
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
