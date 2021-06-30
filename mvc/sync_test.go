package mvc

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	shell "github.com/ipfs/go-ipfs-api"
	_ "github.com/mattn/go-sqlite3"
)

//Test Post upload file.
func TestPostFormDataPublicgatewayFile(t *testing.T) {
	sugar.InitLogger()
	hash, err := PostFormDataPublicgatewayFile("/Users/apple/winter/D-cloud/sugar/", "remote")
	fmt.Println("hash=", hash)
	fmt.Println("err=", err)

}

func TestResolverIpnsAddress(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{name: "", want: "", wantErr: true},
	}
	for _, tt := range tests {
		sugar.InitLogger()

		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolverIpnsAddress()

			t.Log("got:=", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolverIpnsAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResolverIpnsAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDiffBetween(t *testing.T) {
	sugar.InitLogger()
	v1 := ""
	v2 := "QmeFiRN6Vreu2EvECv25yw75GqEdsGw1QF45w4bqGcHYF5_QmeVsha4XkdayJjpzD6ykJR6iXmF3VytyxE18sYCSCkgBR"
	d := FindDiffBetween(v1, v2)
	fmt.Println("d=", d)
}

// 创建 local 文件
func TestLocalNonexistent(t *testing.T) {

	sugar.InitLogger()

	err := LocalNonexistent("/Users/apple/winter/offline/")
	if err != nil {
		fmt.Println("err:=", err)
	}

}

func TestReadRemoteAndLocal(t *testing.T) {
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	// d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	path := "/Users/apple/winter/offline/"
	hash := "QmeFiRN6Vreu2EvECv25yw75GqEdsGw1QF45w4bqGcHYF5"

	gotV1, gotV2, err := ReadRemoteAndLocal(path, hash)
	if err != nil {
		fmt.Println("err:=", err)
	}
	fmt.Println("v1:=", gotV1)
	fmt.Println("v1:=", gotV2)
}

func TestLoopGetCidAndExcuteSql(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Info(" ~~~~~~~~~ Capture the panic ~~~~~~~~~~~~Err: ", err)
		} else {
			sugar.Log.Info("~~~~~~~~~~~~~~~   Normal ~~~~~~~~~~~~")
		}
	}()
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	d, err := sql.Open("sqlite3", "/Users/apple/winter/D-cloud/tables/foo.db")
	if err != nil {
		panic(err)
	}
	db := Testdb(d)
	v1 := ""
	v2 := "QmeFiRN6Vreu2EvECv25yw75GqEdsGw1QF45w4bqGcHYF5_QmeVsha4XkdayJjpzD6ykJR6iXmF3VytyxE18sYCSCkgBR"
	diff := FindDiffBetween(v1, v2)
	path := "/Users/apple/winter/offline/"
	fmt.Println("diff:=", diff)
	fmt.Println("db:=", db)
	err = LoopGetCidAndExcuteSql(diff, path, &db)
	if err != nil {
		fmt.Println("err:=", err)
	}
	fmt.Println("=========")

}
func Testdb(sq *sql.DB) Sql {
	return Sql{DB: sq}
}

func TestUploadFile(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Info(" ~~~~~~~~~ Capture the panic ~~~~~~~~~~~~Err: ", err)
		} else {
			sugar.Log.Info("~~~~~~~~~~~~~~~   Normal ~~~~~~~~~~~~")
		}
	}()
	sh = shell.NewShell("127.0.0.1:5001")
	sugar.InitLogger()
	hash := "QmTD5GfbjzznxAxKgn6H4sEUjHsvtXZd39TRgGY7nnVEo3"
	path := "/Users/apple/winter/offline/"

	got, err := UploadFile(path, hash)
	fmt.Println(err)
	fmt.Println(got)

}
