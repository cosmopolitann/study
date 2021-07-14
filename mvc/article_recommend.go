package mvc

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleRecommend(db *Sql, value string) ([]ArticleAboutMeResp, error) {
	sugar.Log.Info("~~~~ Start   ArticleRecommend ~~~~~")
	var art []ArticleAboutMeResp
	var result vo.ArticleRecommendParams
	//unmarshal params.
	sugar.Log.Info("-Marshal params.")
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return art, err
	}
	sugar.Log.Info("--Marshal data :=", result)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return art, err
	}
	// r := (result.PageNum - 1) * 3
	// sugar.Log.Info("--- pageSize := ", result.PageSize)
	// sugar.Log.Info("--- pageNum := ", result.PageNum)
	//rows, err := db.DB.Query("SELECT * FROM article limit ?,?", r,result.PageSize)
	//SELECT * from article as a LEFT JOIN sys_user as b on a.user_id=b.id  LIMIT 0,4;
	//userid:=cla

	// rows, err := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article order by ptime desc LIMIT ?,?;", r, result.PageSize)
	sugar.Log.Info("---- Excute query data.")

	//

	// 获取 count
	rows1, err := db.DB.Query("select count(*) as count from article")
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, err
	}
	var c int

	for rows1.Next() {
		err = rows1.Scan(&c)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
	}
	defer rows1.Close()
	sugar.Log.Info("----总数c: :=", c)

	rand.Seed(time.Now().UnixNano())

	//step2：获取随机数
	num4 := rand.Intn(c) + 0 //[5,10)

	sugar.Log.Info("----随机数 是 :=", num4)

	rows, err := db.DB.Query("SELECT a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),IFNULL(b.peer_id,'0'),IFNULL(b.name,'0'),IFNULL(b.phone,'0'),IFNULL(b.sex,0),IFNULL(b.nickname,'0'),IFNULL(b.img,''),IFNULL(f.is_like,0),(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on f.article_id=a.id LIMIT ?,?", num4-10, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, err
	}

	// 释放锁
	defer rows.Close()
	for rows.Next() {
		var dl ArticleAboutMeResp
		//scan data =>  ArticleAboutMeResp.
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName, &dl.Img, &dl.IsLike, &dl.LikeNum)
		if err != nil {
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
	sugar.Log.Info("~~~~ Start   ArticleRecommend  End ~~~~~")
	return art, nil
}

// limit 10

func ArticleRecommendLimitTenData(db *Sql, value string) ([]ArticleAboutMeResp, error) {
	sugar.Log.Info("~~~~ Start   ArticleRecommend ~~~~~")
	var art []ArticleAboutMeResp

	sugar.Log.Info("-Marshal params.")
	//
	rows, err := db.DB.Query("SELECT a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),IFNULL(b.peer_id,'0'),IFNULL(b.name,'0'),IFNULL(b.phone,'0'),IFNULL(b.sex,0),IFNULL(b.nickname,'0'),IFNULL(b.img,''),IFNULL(f.is_like,0),(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on f.article_id=a.id order by a.ptime DESC LIMIT ?,?", 0, 10)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, err
	}
	// 释放锁
	defer rows.Close()

	for rows.Next() {
		var dl ArticleAboutMeResp
		//scan data =>  ArticleAboutMeResp.
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName, &dl.Img, &dl.IsLike, &dl.LikeNum)
		if err != nil {
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
	sugar.Log.Info("~~~~ Start   ArticleRecommend  End ~~~~~")
	return art, nil
}
