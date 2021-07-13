package mvc

import (
	"encoding/json"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleAboutMe(db *Sql, value string) ([]ArticleAboutMeResp, error) {
	sugar.Log.Info("~~~~~ Start ArticleAboutMe   ~~~~~ ")
	var art []ArticleAboutMeResp
	var result vo.ArticleAboutMeParams
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return art, err
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
	//check token is valid.
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return art, err
	}
	sugar.Log.Info("claim := ", claim)
	userid := claim["UserId"]
	rows, err := db.DB.Query("SELECT a.is_like,b.id,IFNULL(b.user_id,'null'),IFNULL(b.accesstory,'null'),IFNULL(b.accesstory_type,0),IFNULL(b.text,'null'),IFNULL(b.tag,'null'),IFNULL(b.ptime,0),IFNULL(b.play_num,0),IFNULL(b.share_num,0),IFNULL(b.title,'null'),IFNULL(b.thumbnail,'null'),IFNULL(b.file_name,'null'),IFNULL(b.file_size,0),IFNULL(c.img,''),IFNULL(c.name,''),IFNULL(c.nickname,''),IFNULL(c.peer_id,''),IFNULL(c.phone,''),IFNULL(c.sex,0),( SELECT COUNT( * ) FROM article_like AS c WHERE c.article_id = a.article_id ) as sum from article_like as a LEFT JOIN article as b on a.article_id=b.id LEFT JOIN sys_user as c on c.id=b.user_id where a.user_id=? and a.is_like=1 ORDER BY b.ptime LIMIT ?,?", userid, r, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, err
	}
	for rows.Next() {
		var dl ArticleAboutMeResp
		var id interface{}
		err = rows.Scan(&dl.IsLike, &id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.Img, &dl.Name, &dl.NickName, &dl.PeerId, &dl.Phone, &dl.Sex, &dl.LikeNum)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		if id != "" {
			dl.Id = id.(string)
		}
		sugar.Log.Info("Query a entire data is ", dl)
		if dl.UserId == "" {
			dl.UserId = ""
		}
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Query article is Failed.", err)
		return art, err
	}
	sugar.Log.Info("Query article  is successful.")
	sugar.Log.Info("~~~~~  Start ArticleAboutMe   End ~~~~~ ")
	return art, nil

}
