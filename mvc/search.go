package mvc

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

//查询文件列表

func Search(db *Sql, value string) (data []File, e error) {
	sugar.Log.Info("~~~ Start Search  data  ~~~~")

	var s vo.SearchFileParams
	var arrfile []File

	err := json.Unmarshal([]byte(value), &s)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", s)

	// verify token
	claim, b := jwt.JwtVeriyToken(s.Token)
	if !b {
		return arrfile, errors.New(" token is invaild. ")
	}
	var or string
	if s.Order == "" {
		or = "ptime"
	}
	if s.Order == "time" {
		or = "ptime"
	}
	if s.Order == "name" {
		or = "file_name"

	}
	if s.Order == "type" {
		or = "file_type"

	}
	if s.Order == "size" {
		or = "file_size"

	}
	sugar.Log.Info("order type:", or)
	// userid info claim["UserId"].(string)
	userid := claim["UserId"].(string)
	sugar.Log.Info("claim := ", claim)
	sugar.Log.Info("UserId := ", userid)
	// sql.
	sql := "select id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') from cloud_file where user_id= ? and file_name like'%" + s.Content + "%'" + " order by " + or
	rows, err := db.DB.Query(sql, userid)
	if err != nil {
		sugar.Log.Error("Search data is failed.Err is ", err)
		return arrfile, err
	}
	// 释放锁
	defer rows.Close()
	//scan data.
	for rows.Next() {
		var dl File
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.FileName, &dl.ParentId, &dl.Ptime, &dl.FileCid, &dl.FileSize, &dl.FileType, &dl.IsFolder, &dl.Thumbnail)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("Search a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("Search all data is ", arrfile)
	sugar.Log.Info("~~~ Search  data  End~~~~")
	return arrfile, nil
}

// 文章查询

func ARticleSearch(db *Sql, value string) (data []ArticleAboutMeResp, e error) {
	sugar.Log.Info("~~~ Start   ARticleSearch  data  ~~~~")
	var s vo.ArticleSearchParams
	var arrfile []ArticleAboutMeResp
	//marshal params.
	err := json.Unmarshal([]byte(value), &s)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is :  ", s)
	r := (s.PageNum - 1) * 3
	str := strconv.FormatInt(r, 10)
	pageSize := strconv.FormatInt(s.PageSize, 10)
	// sql := "select id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article where title like'%" + s.Title + "%' limit " + str + "," + pageSize
	//select IFNULL(b.id,''),IFNULL(b.peer_id,''),IFNULL(b.name,''),IFNULL(b.phone,''),IFNULL(b.sex,0),IFNULL(b.ptime,strftime( '%s', 'now')),IFNULL(b.utime,strftime( '%s', 'now')), a.id,IFNULL(a.user_id,''),IFNULL(a.accesstory,''),IFNULL(a.accesstory_type,0),IFNULL(a.text,''),IFNULL(a.tag,''),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,''),IFNULL(a.thumbnail,''),IFNULL(a.file_name,''),IFNULL(a.file_size,0) from article as a LEFT JOIN sys_user as b on a.user_id=b.id where title like '%我%' limit 0,10
	// sql := "select IFNULL(b.peer_id,''),IFNULL(b.name,''),IFNULL(b.phone,''),IFNULL(b.sex,0),IFNULL(b.nickname,0),IFNULL(b.img,0),a.id,IFNULL(a.user_id,''),IFNULL(a.accesstory,''),IFNULL(a.accesstory_type,0),IFNULL(a.text,''),IFNULL(a.tag,''),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,''),IFNULL(a.thumbnail,''),IFNULL(a.file_name,''),IFNULL(a.file_size,0) from article as a LEFT JOIN sys_user as b on a.user_id=b.id where title like'%" + s.Title + "%' limit " + str + "," + pageSize
	sql := "select IFNULL(f.is_like,0),IFNULL(b.peer_id,''),IFNULL(b.name,''),IFNULL(b.phone,''),IFNULL(b.sex,0),IFNULL(b.nickname,0),IFNULL(b.img,0),a.id,IFNULL(a.user_id,''),IFNULL(a.accesstory,''),IFNULL(a.accesstory_type,0),IFNULL(a.text,''),IFNULL(a.tag,''),IFNULL(a.ptime,0),IFNULL(a.play_num,0),IFNULL(a.share_num,0),IFNULL(a.title,''),IFNULL(a.thumbnail,''),IFNULL(a.file_name,''),IFNULL(a.file_size,0),(SELECT count(*) FROM article_like AS d  WHERE d.article_id = a.id AND d.is_like = 1 ) AS likeNum from article as a LEFT JOIN sys_user as b on a.user_id=b.id LEFT JOIN article_like as f on f.article_id=a.id where title like'%" + s.Title + "%' limit " + str + "," + pageSize

	rows, err := db.DB.Query(sql)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		var dl ArticleAboutMeResp
		err = rows.Scan(&dl.IsLike, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.NickName, &dl.Img, &dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize, &dl.LikeNum)
		if err != nil {
			sugar.Log.Error("ARticleSearch scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("ARticleSearch a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("ARticleSearch all data is ", arrfile)
	sugar.Log.Info("~~~   ARticleSearch  data  End~~~~")
	return arrfile, nil
}

//
// 文章查询

func ArticleSearchCagetory(db *Sql, value string) (data []Article, e error) {
	sugar.Log.Info("~~~ Start   ArticleSearchCagetory  data  ~~~~")
	var s vo.ArticleSearchCategoryParams
	var arrfile []Article
	//marshal params.
	err := json.Unmarshal([]byte(value), &s)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is :  ", s)
	r := (s.PageNum - 1) * 3
	str := strconv.FormatInt(r, 10)
	pageSize := strconv.FormatInt(s.PageSize, 10)
	AccesstoryType := strconv.FormatInt(s.AccesstoryType, 10)

	sql := "select id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article where accesstory_type=" + AccesstoryType + " limit " + str + "," + pageSize
	rows, err := db.DB.Query(sql)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return arrfile, errors.New("查询下载列表信息失败")
	}
	defer rows.Close()
	for rows.Next() {
		var dl Article
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("ARticleSearch scan data is failed.The err is ", err)
			return arrfile, err
		}
		sugar.Log.Info("ARticleSearch a entire data is ", dl)
		arrfile = append(arrfile, dl)
	}
	sugar.Log.Info("ARticleSearch all data is ", arrfile)
	sugar.Log.Info("~~~   ARticleSearch  data  End~~~~")
	return arrfile, nil
}
