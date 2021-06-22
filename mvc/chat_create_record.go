package mvc

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatCreateRecord(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) (vo.ChatRecordInfo, error) {

	// 接收参数
	var msg vo.ChatAddRecordParams

	// 返回参数
	var ret vo.ChatRecordInfo

	sugar.Log.Info("Request Param:", value)

	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return ret, err
	}
	sugar.Log.Info("Marshal data is  ", msg)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return ret, errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	userId := claim["UserId"].(string)

	if userId != msg.FromId {
		sugar.Log.Error("token is not msg.from_id")
		return ret, errors.New("token is not msg.from_id")
	}

	ret.Id = genRecordID(msg.FromId, msg.ToId)
	ret.Name = ""
	ret.FromId = msg.FromId
	ret.Toid = msg.ToId
	ret.LastMsg = ""
	ret.Ptime = time.Now().Unix()

	// 检查是否存在
	err = db.DB.QueryRow("SELECT id, from_id, to_id, ptime, last_msg FROM chat_record WHERE id = ?", ret.Id).Scan(&ret.Id, &ret.FromId, &ret.Toid, &ret.Ptime, &ret.LastMsg)
	if err != nil && err != sql.ErrNoRows {
		sugar.Log.Error("Query chat_record failed.Err is", err)
		return ret, err
	}

	switch err {
	case sql.ErrNoRows:
		// no room
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) VALUES (?, ?, ?, ?, ?, ?)", ret.Id, ret.Name, ret.FromId, ret.Toid, ret.Ptime, ret.LastMsg)
		if err != nil {
			sugar.Log.Error("INSERT INTO chat_record is Failed.", err)
			return ret, err
		}

		_, err = res.LastInsertId()
		if err != nil {
			sugar.Log.Error("INSERT INTO chat_record is Failed2.", err)
			return ret, err
		}

	case nil:
		// exists do nothing
	default:
		// error
		return ret, err
	}

	// 查询对方信息
	err = db.DB.QueryRow("SELECT peer_id, name, phone, sex, nickname, img FROM sys_user WHERE id = ?", msg.ToId).Scan(&ret.PeerId, &ret.UserName, &ret.Phone, &ret.Sex, &ret.NickName, &ret.Img)
	if err != nil {
		if err == sql.ErrNoRows {
			sugar.Log.Warn("not found peer info, so set empty")
		} else {
			sugar.Log.Error("Query Peer User Failed. Err:", err)
			return ret, err
		}
	}

	swapMsg := vo.ChatSwapRecordParams{
		Id:      ret.Id,
		Name:    ret.Name,
		Img:     "",
		FromId:  ret.FromId,
		ToId:    ret.Toid,
		Ptime:   ret.Ptime,
		LastMsg: ret.LastMsg,
		Token:   "",
	}

	msgBytes, err := json.Marshal(map[string]interface{}{
		"type": vo.MSG_TYPE_RECORD,
		"data": swapMsg,
	})
	if err != nil {
		sugar.Log.Error("Message record marshal failed.Err is", err)
		return ret, err
	}

	sugar.Log.Info("publish data: ", string(msgBytes))

	ipfsTopic, ok := TopicJoin.Load(vo.CHAT_MSG_SWAP_TOPIC)
	if !ok {
		ipfsTopic, err = ipfsNode.PubSub.Join(vo.CHAT_MSG_SWAP_TOPIC)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return ret, err
		}

		TopicJoin.Store(vo.CHAT_MSG_SWAP_TOPIC, ipfsTopic)
	}

	ctx := context.Background()

	err = ipfsTopic.Publish(ctx, msgBytes)
	if err != nil {
		sugar.Log.Error("publish failed.", err)
		return ret, err
	}

	sugar.Log.Info("publish success")

	return ret, nil
}

func genRecordID(fromID, toID string) string {
	if fromID <= toID {
		return fromID + "_" + toID
	} else {
		return toID + "_" + fromID
	}

}
