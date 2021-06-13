package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ArticleCategory(db *Sql, value string)([]vo.ArticleResp, error) {
	var art []vo.ArticleResp
	var result vo.ArticleCategoryParams
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return art,errors.New("解析错误")
	}
	sugar.Log.Info("Marshal data is  ", result)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return art,err
	}
	sugar.Log.Error("Marshal data is  result := ", result)
	r:=(result.PageNum-1)*3
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	//rows, err := db.DB.Query("SELECT * FROM article limit ?,?", r,result.PageSize)
	//SELECT * from article as a LEFT JOIN sys_user as b on a.user_id=b.id  LIMIT 0,4;
	rows, err := db.DB.Query("SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname from article as a LEFT JOIN sys_user as b on a.user_id=b.id  LIMIT ?,?;", r,result.PageSize)

	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var dl vo.ArticleResp
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType,&dl.Text, &dl.Tag, &dl.Ptime ,&dl.ShareNum,&dl.PlayNum,&dl.Title,&dl.Thumbnail,&dl.PeerId,&dl.Name,&dl.Phone,&dl.Sex,&dl.NickName)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		art = append(art, dl)
	}
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return art,err
	}
	sugar.Log.Info("Query article  is successful.")

	return art,nil

}