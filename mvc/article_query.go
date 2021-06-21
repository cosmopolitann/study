package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//查询文件列表

func ArticleQuery(db *Sql, value string) (data vo.ArticleResp, e error) {
	//var dl Article
	var dl vo.ArticleResp
	var list vo.ArticleQueryParams
	err := json.Unmarshal([]byte(value), &list)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return dl, err
	}
	sugar.Log.Info("Marshal data is  ", list)
	// 查询
	//rows, err := db.DB.Query("select * from article where id=?", list.Id)

	//var islike interface{}

	rows, err := db.DB.Query("SELECT a.*,b.peer_id,b.name,b.phone,b.sex,b.nickname from article as a LEFT JOIN sys_user as b on a.user_id=b.id where a.user_id=(select c.user_id from article as c where id =?) and a.id=? order by ptime desc;", list.Id, list.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return dl, errors.New("查询下载列表信息失败")
	}
	for rows.Next() {
		var peerId interface{}
		var name interface{}
		var phone interface{}
		var sex interface{}
		var NickName interface{}
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.ShareNum, &dl.PlayNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &peerId, &name, &phone, &sex, &NickName)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return dl, err
		}
		var k = ""
		if peerId == nil {
			dl.PeerId = k
			dl.PeerId = k
			dl.Name = k
			dl.Phone = k
			dl.Sex = 0
			dl.NickName = k
			dl.LikeNum = 0
		} else {
			dl.PeerId = peerId.(string)
			dl.Name = name.(string)
			dl.Phone = phone.(string)
			dl.Sex = sex.(int64)
			dl.NickName = NickName.(string)
			//dl.IsLike=islike.(int64)

		}
	}

	sugar.Log.Info("Query all data is ", dl)
	return dl, nil

}
