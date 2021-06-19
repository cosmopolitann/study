package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleAboutMe(db *Sql, value string) ([]Article, error) {
	var art []Article
	var result vo.ArticleAboutMeParams
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return art, errors.New("解析错误")
	}
	sugar.Log.Info("Marshal data is  ", result)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return art, err
	}
	sugar.Log.Error("Marshal data is  result := ", result)
	r := (result.PageNum - 1) * result.PageSize
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	//rows, err := db.DB.Query("SELECT * FROM article limit ?,?", r,result.PageSize)
	//SELECT * from article as a LEFT JOIN sys_user as b on a.user_id=b.id  LIMIT 0,4;
	//userid:=cla

	//SELECT * from article_like as a where user_id='409330202166956032' and is_like=1;

	//token
	//验证token 是否满足条件
	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return art, err
	}
	sugar.Log.Info("claim := ", claim)
	userid := claim["UserId"]
	rows, err := db.DB.Query("SELECT b.* from article_like as a LEFT JOIN article as b on a.article_id=b.id WHERE a.is_like=1 and a.user_id=? ORDER BY ptime LIMIT ?,?", userid, r, result.PageSize)

	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl Article
		var id interface{}
		//var peerId interface{}
		//var name interface{}
		//var phone interface{}
		//var sex interface{}
		//var NickName interface{}
		//var islike interface{}PlayNum
		err = rows.Scan(&id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		if id != "" {
			dl.Id = id.(string)
		}

		sugar.Log.Info("Query a entire data is ", dl)
		if dl.UserId == "" {
			dl.UserId = "anonymity"
		}
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return art, err
	}
	sugar.Log.Info("Query article  is successful.")

	return art, nil

}
