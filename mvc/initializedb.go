package mvc

import (
	"database/sql"

	"github.com/cosmopolitann/clouddb/sugar"
)

type Sql struct {
	DB *sql.DB
}

type NewTestNode struct {
	db Sql
}

// dbPath  - /path/to/foo.db
// logPath - /path/to/log.txt
// env     - development、test、production
func NTestNode(dbPath, logPath, env string) *NewTestNode {
	sugar.InitLogger1(logPath, env)
	sugar.Log.Info("~~~~  Connecting to the sqlite3 database. ~~~~")
	sql := Newdb(dbPath)

	return &NewTestNode{db: sql}
}

func (n *NewTestNode) Add() error {
	//
	err := n.db.Ping()
	if err != nil {
		sugar.Log.Error("Open db is failed. Err: ", err)
	}
	return err
}

func (n *NewTestNode) UserLogin(value string) string {
	//
	data := n.db.UserLogin(value)
	return data
}

func Newdb(path string) Sql {
	return Sql{DB: InitDB(path)}
}

func InitDB(path string) *sql.DB {
	//
	//mvc, err := sql.Open("sqlite3", path)
	if path == "" {
		path = "../tables/foo.db"
	}
	sugar.Log.Info(" Db path := ", path)
	sugar.Log.Info("Start Open Sqlite3 Database.")
	// ?_fk=1 启动外键
	db, err := sql.Open("sqlite3", path+"?_fk=1")
	checkErr(err)
	sugar.Log.Info("Open Sqlite3 is ok.")
	sugar.Log.Info("Db value is ", db)

	return db
}
func checkErr(err error) {
	if err != nil {
		sugar.Log.Error("The connection to the database failed.")
		panic(err)
	}
}
