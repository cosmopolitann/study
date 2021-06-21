package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleCategory(db *Sql, value string) ([]vo.ArticleResp, error) {
	var art []vo.ArticleResp
	var result vo.ArticleCategoryParams
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
	r1 := result.PageSize

	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	//rows, err := db.DB.Query("SELECT * FROM article limit ?,?", r,result.PageSize)
	//SELECT * from article as a LEFT JOIN sys_user as b on a.user_id=b.id  LIMIT 0,4;
	//userid:=cla
	//rows, err := db.DB.Query("SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname ,(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id WHERE a.accesstory_type =?  ORDER BY ptime desc LIMIT ?,?", result.AccesstoryType, r, r1)
	//SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname ,(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum ,f.is_like FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on a.id=f.article_id WHERE a.accesstory_type =2
	rows, err := db.DB.Query("SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname ,(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum ,IFNULL(f.is_like,0) FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id LEFT JOIN article_like as f on a.id=f.article_id WHERE a.accesstory_type =? ORDER BY ptime desc LIMIT ?,?", result.AccesstoryType, r, r1)

	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl vo.ArticleResp
		var peerId interface{}
		var name interface{}
		var phone interface{}
		var sex interface{}
		var NickName interface{}
		//var likeNum int64
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &peerId, &name, &phone, &sex, &NickName, &dl.LikeNum,&dl.Islike)
		if err != nil { //PlayNum
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		sugar.Log.Info("数据  dl=== ", dl)

		var k = ""
		if peerId == nil  {
			dl.PeerId = k
			dl.PeerId = k
			dl.Name = k
			dl.Phone = k
			dl.Sex = 0
			dl.NickName = k
		} else {
			dl.PeerId = peerId.(string)
			dl.Name = name.(string)
			dl.Phone = phone.(string)
			dl.Sex = sex.(int64)
			dl.NickName = NickName.(string)
		}
		//dl.PeerId = peerId.(string)
		//dl.Name = name.(string)
		//dl.Phone = phone.(string)
		//dl.Sex = sex.(int64)
		//dl.NickName = NickName.(string)
		sugar.Log.Info("345 ")

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
	sugar.Log.Info("查出来的结果：len ",len(art))

	return art, nil

}
