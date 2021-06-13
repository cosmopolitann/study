package chat
import (
	"github.com/cosmopolitann/clouddb/sugar"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestChatMsgWithDraw(t *testing.T) {
	sugar.InitLogger()
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	//The path is default.
	sugar.Log.Info("Start Open Sqlite3 Database.")
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err!=nil{
		panic(err)
	}
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ",d)
	e:=d.Ping()
	fmt.Println(" Ping is failed,err:=",e)
	ss:= Testdb(d)
	value:=`{"id":"408583164135477248"}
`
	resp:= ss.ChatMsgWithDraw(value)
	t.Log("获取返回的数据 :=  ",resp)

}