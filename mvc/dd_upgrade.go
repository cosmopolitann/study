package mvc

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func DbUpgrade(db *Sql, dbv string) (string, error) {
	// 数据库升级
	//查询当前目录是否存在 version 文件
	//如果存在 查询 版本信息  和 当前的比对

	// //读文件
	// ex := CheckFileIsExist(path + vo.DBversion)
	// if !ex {
	// 	// 文件不存在 就返回出去
	// 	return errors.New("文件不存在")
	// }
	// //读文件信息
	// data, err := ReadContent(path)
	// if err != nil {
	// 	return err
	// }
	// //
	// fmt.Println("data:=", data)
	// fmt.Println("len(data):= ", len(data))
	// if len(data) <= 1 {
	// 	return errors.New("版本数据不对")
	// }
	//读出本地版本
	if dbv == "" {
		sugar.Log.Info("数据库版本字符串为空，不符合")
		return "", errors.New("数据库版本字符串为空，不符合")
	}
	loaclVersion := vo.Version
	lv, _ := strconv.Atoi(loaclVersion)
	uv, _ := strconv.Atoi(dbv)
	if uv < 1 {
		sugar.Log.Info("要更新的版本错误 不更新")
		return "", errors.New("要更新的版本错误 不更新")
	}
	if lv == uv {
		sugar.Log.Info("版本信息一致  不更新")
		return loaclVersion, nil
	}
	//循环遍历 执行sql语句
	if uv > lv {
		sugar.Log.Info("当前版本大于最新版本，不符合")
		return dbv, nil
	}
	for i := uv; i < lv; i++ {
		//执行sql 语句
		result := vo.UpgradeSql[i]
		for _, v := range result {
			result, err := db.DB.Exec(v)
			sugar.Log.Info("执行的sql 语句:= ", v)

			if err != nil {
				sugar.Log.Error("err:= ", err)
				continue
			}
			sugar.Log.Info("result:= ", result)
		}
	}
	// //更新文件版本号
	// local_f, err1 := os.OpenFile(path+vo.DBversion, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //open file.
	// if err1 != nil {
	// 	sugar.Log.Error(" Open file is failed.Err:", err1)
	// 	return err
	// }

	// _, err = local_f.WriteString(string("3"))
	// if err != nil {
	// 	sugar.Log.Error(" Write remote content to this local file is failed.Err: ", err)
	// }
	sugar.Log.Info("~~~  执行完成  ~~~")
	return loaclVersion, nil
}
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ReadContent(path string) (string, error) {
	data, err := ioutil.ReadFile(path + vo.DBversion)

	if err != nil {
		fmt.Println("File reading error", err)
		return "", err
	}
	fmt.Println("Contents of file:", string(data))
	return string(data), nil
}
