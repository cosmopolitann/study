package mvc

import (
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatMsgList(db *Sql, value string) ([]ChatMsg, error) {
	var art []ChatMsg
	var result vo.ChatMsgListParams

	sugar.Log.Debug("Request Param:", value)

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

	sugar.Log.Info("claim := ", claim)
	sugar.Log.Info("Marshal data is  result := ", result)
	r := (result.PageNum - 1) * result.PageSize

	//这里 要修改   加上 where  参数 判断

	//todo
	rows, err := db.DB.Query("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id, send_state, send_fail FROM chat_msg where record_id =? order by ptime desc limit ?,? ", result.RecordId, r, result.PageSize)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return art, errors.New("查询下载列表信息失败")
	}

	// 释放锁
	defer rows.Close()

	for rows.Next() {
		var dl ChatMsg
		err = rows.Scan(&dl.Id, &dl.ContentType, &dl.Content, &dl.FromId, &dl.ToId, &dl.Ptime, &dl.IsWithdraw, &dl.IsRead, &dl.RecordId, &dl.SendState, &dl.SendFail)
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
