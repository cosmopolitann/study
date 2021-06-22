package mvc

import (
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	"strconv"
	"time"
)

func DownLoadFile(db *Sql, value string) (e error) {
	var d vo.TransferAdd
	id := utils.SnowId()
	err := json.Unmarshal([]byte(value), &d)
	if err != nil {
		return err
	}

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(d.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)

	//t := time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO cloud_transfer (id,user_id,file_name,ptime,file_cid,file_size,down_path,file_type,transfer_type,upload_parent_id,upload_file_id) values(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into cloud_down table is failed.", err)
		return errors.New("插入cloud_down 表 数据失败")
	}

	sid := strconv.FormatInt(id, 10)
	res, err := stmt.Exec(sid, claim["UserId"].(string), d.FileName, t, d.FileCid, d.FileSize, d.FilePath, d.FileType, d.TransferType, d.UploadParentId, d.UploadFileId)

	if err != nil {
		sugar.Log.Error("Insert into cloud_down  is Failed.", err)
		return err
	}
	c, _ := res.RowsAffected()
	if c == 0 {
		sugar.Log.Error("Insert into cloud_down  is Failed.", err)
		return errors.New("插入cloud_down表数据失败")
	}
	return nil
}
