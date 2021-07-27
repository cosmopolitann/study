package mvc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

const (
	RECORD_TYPE_USER   = "user"
	RECORD_TYPE_RECORD = "record"
)

func ChatMsgList(db *Sql, value string) ([]ChatMsg, error) {
	var art []ChatMsg
	var result vo.ChatMsgListParams
	err := json.Unmarshal([]byte(value), &result)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", result)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(result.Token)
	if !b {
		return art, errors.New("token 失效")
	}
	userId := claim["id"].(string)

	sugar.Log.Info("claim := ", claim)
	sugar.Log.Info("Marshal data is  result := ", result)
	r := (result.PageNum - 1) * result.PageSize
	sugar.Log.Info("pageSize := ", result.PageSize)
	sugar.Log.Info("pageNum := ", result.PageNum)
	sugar.Log.Info("r := ", r)
	sugar.Log.Info("recordType ==== := ", result.RecordType)
	sugar.Log.Info("recordId ==== := ", result.RecordId)

	recordId := result.RecordId
	if result.RecordType == RECORD_TYPE_USER {
		toUserId := result.RecordId
		recordId = genRecordID(userId, toUserId)
		var rid string
		err := db.DB.QueryRow("SELECT id FROM chat_record where id = ?", recordId).Scan(&rid)
		if err != nil && err != sql.ErrNoRows {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return art, errors.New("查询下载列表信息失败")
		}

		if rid == "" {
			// no room
			res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) VALUES (?, ?, ?, ?, ?, ?)", recordId, "", userId, toUserId, time.Now().Unix(), "")
			if err != nil {
				sugar.Log.Error("INSERT INTO chat_record is Failed.", err)
				return art, err
			}

			_, err = res.LastInsertId()
			if err != nil {
				sugar.Log.Error("INSERT INTO chat_record is Failed2.", err)
				return art, err
			}
		}
	}

	//这里 要修改   加上 where  参数 判断

	//todo
	rows, err := db.DB.Query("SELECT * FROM chat_msg where record_id =? order by ptime desc limit ?,? ", recordId, r, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}

	// 释放锁
	defer rows.Close()

	for rows.Next() {
		var dl ChatMsg
		err = rows.Scan(&dl.Id, &dl.ContentType, &dl.Content, &dl.FromId, &dl.ToId, &dl.Ptime, &dl.IsWithdraw, &dl.IsRead, &dl.RecordId)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return art, err
		}
		sugar.Log.Info("Query a entire data is ", dl)
		art = append(art, dl)
	}

	for i, j := 0, len(art)-1; i < j; i, j = i+1, j-1 {
		art[i], art[j] = art[j], art[i]
	}

	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return art, err
	}

	sugar.Log.Info("Insert into article  is successful.")
	return art, nil

}
