package sql

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	_ "github.com/mattn/go-sqlite3"
)

func TestFileStatExist(t *testing.T) {

	b := Check("/Users/apple/winter/clouddb/vo/db-version.txt")
	fmt.Println("b:=", b)
	if !b {
		fmt.Println("文件 不存在")
	}
	fmt.Println("11111")

	data, err := ioutil.ReadFile("/Users/apple/winter/clouddb/vo/db-version.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))

}
func Check(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func ReadFileConenet(path string) (string, error) {
	fmt.Println("路径：=", path+vo.DBversion)
	file, err := os.Open(path + vo.DBversion)
	if err != nil {
		fmt.Println("err:=", err)
	}
	defer file.Close()
	f, _ := ioutil.ReadAll(file)
	sugar.Log.Info("查出当前数据库版本信息 version: ", string(f))
	return string(f), nil
}
