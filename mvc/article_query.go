package mvc

import (
	"database/sql"
	"encoding/json"

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
	err = db.DB.QueryRow("SELECT id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0),IFNULL(b.peer_id,'0'),IFNULL(b.name,'0'),IFNULL(b.phone,'0'),IFNULL(b.sex,0),IFNULL(b.nickname,'0'),IFNULL(b.img,'') from article as a LEFT JOIN sys_user as b on a.user_id=b.id where a.user_id=(select c.user_id from article as c where id =?) and a.id=? order by ptime desc;", list.Id, list.Id).
		Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.ShareNum, &dl.PlayNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName, &dl.Img)
	if err != nil && err != sql.ErrNoRows {
		sugar.Log.Error("Query chat_record failed.Err is", err)
		return dl, err
	}
	sugar.Log.Info("Query all data is ", dl)
	return dl, nil
}
