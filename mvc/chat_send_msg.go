package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func ChatSendMsg(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) (ChatMsg, error) {

	// 接收参数
	var msg vo.ChatSendMsgParams

	// 返回参数
	var ret ChatMsg

	sugar.Log.Debug("Request Param:", value)

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
	userId := claim["id"].(string)

	if userId != msg.FromId {
		sugar.Log.Error("token is not msg.from_id")
		return ret, errors.New("token is not msg.from_id")
	}

	ret.Id = strconv.FormatInt(utils.SnowId(), 10)
	ret.ContentType = msg.ContentType
	ret.Content = msg.Content
	ret.FromId = msg.FromId
	ret.ToId = msg.ToId
	ret.Ptime = time.Now().Unix()
	ret.IsWithdraw = 0
	ret.IsRead = 0
	ret.RecordId = msg.RecordId
	ret.SendState = vo.MSG_STATE_SENDING
	ret.SendFail = ""

	res, err := db.DB.Exec(
		"INSERT INTO chat_msg (id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		ret.Id, ret.ContentType, ret.Content, ret.FromId, ret.ToId, ret.Ptime, ret.IsWithdraw, ret.IsRead, ret.RecordId)
	if err != nil {
		sugar.Log.Error("INSERT INTO chat_msg is Failed.", err)
		return ret, err
	}

	_, err = res.LastInsertId()
	if err != nil {
		sugar.Log.Error("INSERT INTO chat_msg is Failed2.", err)
		return ret, err
	}

	// update last msg
	_, err = db.DB.Exec("UPDATE chat_record SET last_msg = ?, ptime = ? WHERE id = ?", ret.Content, ret.Ptime, ret.RecordId)
	if err != nil {
		sugar.Log.Error("UPDATE chat_record .Err is", err)
		return ret, err
	}

	swapMsg := vo.ChatSwapMsgParams{
		Id:          ret.Id,
		RecordId:    ret.RecordId,
		ContentType: ret.ContentType,
		Content:     ret.Content,
		FromId:      ret.FromId,
		ToId:        ret.ToId,
		IsWithdraw:  ret.IsWithdraw,
		IsRead:      ret.IsRead,
		Ptime:       ret.Ptime,
		Token:       "",
	}

	go func() {
		var sendState int64
		var sendFail string

		tryTimes := 0
		maxTimes := 3

		for {
			err = chatSendMsg(ipfsNode, swapMsg)
			if err != nil {
				sugar.Log.Errorf("send chat msg failed. msgid: %s, err: %v", ret.Id, err)
				return
			}

			<-time.After(10 * time.Second)

			tryTimes++
			err := db.DB.QueryRow("select send_state from chat_msg where id = ?", ret.Id).Scan(&sendState)
			if err != nil {
				sendState = -1
				sendFail = err.Error()
				sugar.Log.Error("select send_state from chat_msg err: ", err)
				break
			}

			if sendState != 0 {
				break
			} else if tryTimes >= maxTimes {
				sendState = -1
				sendFail = "failed"
				sugar.Log.Warnf("try over max times %d", maxTimes)
				break
			}
		}

		if sendState == -1 {
			_, err := db.DB.Exec("update chat_msg set send_state = ?, send_fail = ? where id = ?", sendState, sendFail, ret.Id)
			if err != nil {
				sugar.Log.Error("update chat_msg send_state fail", err)
			}

			sugar.Log.Warn("chat send msg failed")
		} else {
			sugar.Log.Warn("chat send msg success")
		}
	}()

	// 发布消息
	return ret, nil
}

func chatSendMsg(ipfsNode *ipfsCore.IpfsNode, swapMsg vo.ChatSwapMsgParams) error {

	var err error

	msgPacket := vo.ChatPacketParams{
		Type: vo.MSG_TYPE_NEW,
		From: ipfsNode.Identity.String(),
		Data: swapMsg,
	}

	msgTopicKey := getRecvTopic(swapMsg.ToId)

	ipfsTopic, ok := TopicJoin.Load(msgTopicKey)
	if !ok {
		ipfsTopic, err = ipfsNode.PubSub.Join(msgTopicKey)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}

		TopicJoin.Store(msgTopicKey, ipfsTopic)
	}

	msgBytes, err := json.Marshal(msgPacket)
	if err != nil {
		sugar.Log.Error("marshal send msg failed.", err)
		return err
	}

	err = ipfsTopic.Publish(context.Background(), msgBytes)
	if err != nil {
		sugar.Log.Error("ChatSendMsg failed.", err)
		return err
	}

	sugar.Log.Debugf("ChatSendMsg topic: %s, data: %v", msgTopicKey, msgPacket)

	return nil
}

// 获取广播topic
func getRecvTopic(toUserId string) string {
	// return vo.CHAT_MSG_SWAP_TOPIC + toUserId
	return vo.CHAT_MSG_SWAP_TOPIC
}
