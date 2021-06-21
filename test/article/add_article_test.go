package article

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/mvc"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"testing"
	"time"
)

func TestAddArticle(t *testing.T) {
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
	fmt.Println(" Ping is failed,err:=", e)
	value:=`{
   "id": "325707698052770816",
   "accesstory": "QMabcdefghijk96",
   "text":"正文开始内容",
   "accesstory_type": 0,
   "tag": "标签31",
   "ptime": "1623955566",
  "play_num": 0,
   "share_num": 0,
   "title": "title65",
   "user_id": "323733228975432704",
   "thumbnail": "thumbnail5",
   "file_name": "",
   "file_size": ""
}`
	var art vo.ArticleAddParams
	err = json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)
	id := utils.SnowId()
	t1:=time.Now().Unix()
	stmt, err := d.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t1, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
	}
	l, _ := res.RowsAffected()
	if l == 0 {
	}
	resp := (value)
	fmt.Println("这是返回的数据 =", resp)

}
func Testdb2(sq *sql.DB) mvc.Sql {
	return mvc.Sql{DB: sq}
}
