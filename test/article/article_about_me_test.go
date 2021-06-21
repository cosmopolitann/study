package article

import (
	"database/sql"
	"github.com/cosmopolitann/clouddb/sugar"
	"testing"
)

func TestArticleAboutMe(t *testing.T) {


	t.Run("insert article",func(t *testing.T) {
		sugar.InitLogger()
		sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
		//The path is default.
		sugar.Log.Info("Start Open Sqlite3 Database.")
		d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
		if err != nil {
			panic(err)
		}
		sugar.Log.Info("Open Sqlite3 is ok.")
		sugar.Log.Info("Db value is ", d)
		e := d.Ping()
		t.Log(" Ping is failed,err:=", e)
		value:=`{
   "id": "325707698052770816",
   "accesstory": "QMabcdefghijk96",
   "text":"古剑奇谭",
   "accesstoryType": 2,
   "tag": "标签01",
   "ptime": "1623955566",
  "playNum": 0,
   "shareNum": 0,
   "title": "title65",
   "userId": "409330202166956032",
   "thumbnail": "thumbnail5",
   "fileName": "",
   "fileSize": ""
}`

		ss := Testdb(d)
		resp := ss.ArticleAddTest(value)
		t.Log("result:=", resp)
	})



	t.Run("article_about_me", func(t *testing.T) {
		sugar.InitLogger()
		sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
		//The path is default.
		sugar.Log.Info("Start Open Sqlite3 Database.")
		d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
		if err != nil {
			panic(err)
		}
		sugar.Log.Info("Open Sqlite3 is ok.")
		sugar.Log.Info("Db value is ", d)
		e := d.Ping()
		sugar.Log.Info(" Ping is failed,err:= ", e)
		ss := Testdb(d)
		// request json  params
		// test 1
		value := `{"pageSize":3,"pageNum":1,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI"}
`
		t.Log("request value :=", value)
		resp := ss.ArticleAboutMe(value)
		t.Log("result:=", resp)
	})



}
