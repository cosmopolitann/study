package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleRecommend(db *Sql, value string) ([]ArticleAboutMeResp, error) {
	var art []ArticleAboutMeResp
	var result vo.ArticleRecommendParams
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
	r := (result.PageNum - 1) * 3
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	//rows, err := db.DB.Query("SELECT * FROM article limit ?,?", r,result.PageSize)
	//SELECT * from article as a LEFT JOIN sys_user as b on a.user_id=b.id  LIMIT 0,4;
	//userid:=cla

	// rows, err := db.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article order by ptime desc LIMIT ?,?;", r, result.PageSize)
	rows, err := db.DB.Query("SELECT a.id,IFNULL(a.user_id,'null'),IFNULL(a.accesstory,'null'),IFNULL(a.accesstory_type,0),IFNULL(a.text,'null'),IFNULL(a.tag,'null'),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,'null'),IFNULL(a.thumbnail,'null'),IFNULL(a.file_name,'null'),IFNULL(a.file_size,0),IFNULL(b.peer_id,'0'),IFNULL(b.name,'0'),IFNULL(b.phone,'0'),IFNULL(b.sex,0),IFNULL(b.nickname,'0') FROM article AS a LEFT JOIN sys_user AS b ON a.user_id = b.id where play_num>1 or share_num>1 ORDER BY a.ptime desc LIMIT ?,?", r, result.PageSize)

	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl ArticleAboutMeResp

		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
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
