package sync

import (
	"database/sql"
	"encoding/json"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	"log"
	"testing"
)


func TestSyncA(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "../../tables/foo.db")
	if err != nil {
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", d)
	e := d.Ping()
	log.Println(" Ping is failed,err:=", e)
	//ss := Testdb(d)
	//	value:=`{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MDkzMzAyMDIxNjY5NTYwMzIiLCJleHAiOjE2MjU4ODk0NzZ9.OzEFVuB2FcRYurZiii1fpiAqX2KcesfS5arJfVJZQOI","content":"三国"}
	//`
	syncValue := `{"method":"SyncUser","data":{"id":"4324","peerId":"124","name":"20","phone":"1889","sex":"1","nickName":"nick"}}`

	var sc vo.SyncParams
	err = json.Unmarshal([]byte(syncValue), &sc)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	log.Println(" 解析的值 =", sc)
	if sc.Method == "SyncUser" {
		//
		//json 转成 string

		jsonBytes, err := json.Marshal(sc.Data)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("转换为 json 串打印结果:%s", string(jsonBytes))
		//resp := ss.SyncArticle(string(jsonBytes))
		//log.Println("这是返回的数据 =", resp)
	}
	log.Println("方法不匹配")
}
