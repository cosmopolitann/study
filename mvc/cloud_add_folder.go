package mvc

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"github.com/goinggo/mapstructure"
)

func AddFolder(db *Sql, value string) error {
	//add folder
	sugar.Log.Info("---- AddFolder   Method  ---- ")

	var f vo.CloudAddFolderParams
	f1 := ConvertString(value, f)
	err := mapstructure.Decode(f1, &f)
	sugar.Log.Info("---- 开始 解析 参数  ---- 参数：", f)
	sugar.Log.Info("Decode data  is  ", f)
	sugar.Log.Info("参数 ParentId: ", f.ParentId)
	sugar.Log.Info("参数 FileName:", f.FileName)
	sugar.Log.Info("参数 Id:", f.Id)
	sugar.Log.Info("参数 Token:", f.Token)
	sugar.Log.Info("参数 Thumbnail:", f.Thumbnail)

	if err != nil {
		sugar.Log.Error("Decode is failed.", err)
		return errors.New("解析map失败")
	}
	//token
	claim, b := jwt.JwtVeriyToken(f.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("校验 token 成功   Token:", claim["UserId"])

	sugar.Log.Info("claim := ", claim)

	userId := claim["UserId"].(string)

	sugar.Log.Info("  查看 user id  ", userId)

	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功 1")

	e := IsFormat(f)

	if e != nil {
		return err
	}
	sugar.Log.Info("判断文件夹 是否 不满足 格式 成功  2")

	//count,in:= InsertIntoData(db,f,userId.(string))
	//	if in!=nil || count==0{
	//		return errors.New("创建文件夹失败")
	//	}

	sugar.Log.Info("-- 查找是否有 相同名字的文件夹 ---")
	c, _ := FindOneDirIsExist(db, f)
	sugar.Log.Info("--  c == ---", c)

	if c == 0 {
		count, in := InsertIntoData(db, f, userId)
		if in != nil || count == 0 {
			return errors.New(" Create folder is failed. ")
		}
	}

	//-1 代表 文件，名字相同 但 cid 不相同
	if c == 1 {
		//后缀名
		timeUnix := time.Now().Unix()
		timeUnixStr := strconv.FormatInt(timeUnix, 10)
		f.FileName = f.FileName + "_" + timeUnixStr
		count, in := InsertIntoData(db, f, userId)
		if in != nil || count == 0 {
			return errors.New(" Create folder is failed. ")
		}
	}
	sugar.Log.Info("---- AddFolder   Method  End ---- ")
	return nil
}
func InsertIntoData(db *Sql, f vo.CloudAddFolderParams, userId string) (c int64, e error) {
	//insert into
	//snowId
	id := utils.SnowId()
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	sugar.Log.Info("--  开始  插入文件夹  ---")
	sugar.Log.Info("--  参数信息   ParentId---", f.ParentId)
	sugar.Log.Info("--  参数信息   ParentId---", f.FileName)
	sugar.Log.Info("--  参数信息   ---", f)

	stmt, err := db.DB.Prepare("INSERT INTO cloud_file (id,user_id,file_name,parent_id,ptime,file_cid,file_size,file_type,is_folder,thumbnail) values(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into cloud_file table is failed.", err)
		return 0, err
	}
	sid := strconv.FormatInt(id, 10)

	res, err := stmt.Exec(sid, userId, f.FileName, f.ParentId, t, "", 0, 0, 1, f.Thumbnail)
	if err != nil {
		sugar.Log.Error("Insert into file  is Failed.", err)
		return 0, err
	}
	sugar.Log.Info("Insert into file  is successful.")
	l, err := res.RowsAffected()
	if l == 0 {
		return 0, err
	}

	sugar.Log.Info("--  插入文件的 雪花id  = ---", sid)

	return l, nil
}

func IsFormat(f vo.CloudAddFolderParams) error {

	//pId, err := strconv.Atoi(f.ParentId)
	pId, err := strconv.ParseInt(f.ParentId, 10, 64)
	if err != nil {
		return err
	}
	if pId < 0 {
		return errors.New("参数不能为负数")
	}
	if IsEmptyRename(f.FileName) {
		return errors.New("文件夹名称不能为空")
	}
	if IsLenRename(f.FileName) {
		return errors.New("文件夹名称过长")
	}
	if f.FileName[0] == vo.IllegalPoint {
		return errors.New("文件夹不能为包含非法字符")
	}
	if find := strings.Contains(f.FileName, "/"); find {
		return errors.New("文件夹不能为包含非法字符")
	}
	if find := strings.Contains(f.FileName, "\\"); find {
		return errors.New("文件夹不能为包含非法字符")
	}
	if find := strings.Contains(f.FileName, "*"); find {
		return errors.New("文件夹不能为包含非法字符")
	}
	return nil
}

func IsLenRename(rename string) bool {
	if len(rename) > 256 {
		return true
	}
	return false
}
func IsEmptyRename(rename string) bool {
	if len(rename) == 0 {
		return true
	}
	return false
}

func FindOneDirIsExist(mvc *Sql, d vo.CloudAddFolderParams) (int64, error) {
	//查询数据
	var f File
	rows, _ := mvc.DB.Query("SELECT id,IFNULL(user_id,'null'),IFNULL(file_name,'null'),IFNULL(parent_id,0),IFNULL(ptime,0),IFNULL(file_cid,'null'),IFNULL(file_size,0),IFNULL(file_type,0),IFNULL(is_folder,0),IFNULL(thumbnail,'null') FROM cloud_file where file_name=? and parent_id=? and is_folder=?", d.FileName, d.ParentId, 1)

	// 释放锁
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&f.Id, &f.UserId, &f.FileName, &f.ParentId, &f.Ptime, &f.FileCid, &f.FileSize, &f.FileType, &f.IsFolder, &f.Thumbnail)
		if err != nil {
			return 0, err
		}
	}
	if f.Id != "" {
		//如果文件 cid 不相等  返回 0
		return 1, nil
	}
	return 0, nil
}
