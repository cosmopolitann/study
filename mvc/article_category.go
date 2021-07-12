package mvc

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleCategory(db *Sql, value string) ([]vo.ArticleResp, error) {
	sugar.Log.Info("----  ArticleCategory  Method  -----")
	var art []vo.ArticleResp
	var result vo.ArticleCategoryParams
	//unmarshal params info.
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
	//get userid
	sugar.Log.Error("Marshal data is  result := ", result)
	r := (result.PageNum - 1) * result.PageSize
	r1 := result.PageSize
	//rows, err := db.DB.Query("SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname ,(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id WHERE a.accesstory_type =?  ORDER BY ptime desc LIMIT ?,?", result.AccesstoryType, r, r1)
	//SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname ,(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum ,f.is_like FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on a.id=f.article_id WHERE a.accesstory_type =2
	var rows *sql.Rows
	if result.Token == "" {
		//query
		rows, err = db.DB.Query("SELECT a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),IFNULL(b.peer_id,''),IFNULL(b.name,''),IFNULL(b.phone,''),IFNULL(b.sex,0),IFNULL(b.nickname,''),IFNULL(b.img,''),(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum ,IFNULL(f.is_like,0) FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on a.id=f.article_id  WHERE a.accesstory_type =? ORDER BY a.ptime desc LIMIT ?,?", result.AccesstoryType, r, r1)
	} else {
		claim, b := jwt.JwtVeriyToken(result.Token)
		if !b {
			return art, errors.New(" Token is valid. ")
		}
		user_id := claim["UserId"]
		rows, err = db.DB.Query("SELECT a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),IFNULL(b.peer_id,''),IFNULL(b.name,''),IFNULL(b.phone,''),IFNULL(b.sex,0),IFNULL(b.nickname,'') ,IFNULL(b.img,''),(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum ,IFNULL(f.is_like,0) FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on a.id=f.article_id and f.user_id=? WHERE a.accesstory_type =? ORDER BY a.ptime desc LIMIT ?,?", user_id, result.AccesstoryType, r, r1)
	}

	// rows, err := db.DB.Query("SELECT a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),IFNULL(b.peer_id,'0'),IFNULL(b.name,'0'),IFNULL(b.phone,'0'),IFNULL(b.sex,0),IFNULL(b.nickname,'0') ,(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum ,IFNULL(f.is_like,0) FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on a.id=f.article_id and f.user_id=? WHERE a.accesstory_type =? ORDER BY a.ptime desc LIMIT ?,?", user_id, result.AccesstoryType, r, r1)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, err
	}
	for rows.Next() {
		var dl vo.ArticleResp
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName, &dl.Img, &dl.LikeNum, &dl.Islike)
		if err != nil {
			//PlayNum
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		if dl.UserId == "" {
			dl.UserId = ""
		}
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return art, err
	}
	sugar.Log.Info("Query article  is successful.")
	sugar.Log.Info("----  ArticleCategory  Method   End -----")
	return art, nil
}
