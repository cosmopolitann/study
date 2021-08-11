package mvc

import (
	"context"
	bsql "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

var curSub *pubsub.Subscription
var listenUserId string
var listenCancelFunc context.CancelFunc

func ChatListenMsgUpdateUser(ipfsNode *ipfsCore.IpfsNode, token string) error {

	if token == "" {
		listenUserId = ""
		sugar.Log.Info("Anonymous User Listen")
	} else {
		userId, err := parseToken(token)
		if err != nil {
			sugar.Log.Errorf("parseToken failed, token: %s, error: %s \n", token, err.Error())
			return errors.New("token is invaild")
		}

		listenUserId = userId
		sugar.Log.Infof("Named User Listen %s", listenUserId)
	}

	err := updateIpfsTopicSubs(ipfsNode, getRecvTopic(listenUserId))
	if err != nil {
		sugar.Log.Error("subscribe failed")
		return fmt.Errorf("subscribe failed")
	}

	if listenCancelFunc != nil {
		listenCancelFunc()
	}

	return nil
}

func ChatListenMsgBlocked(ipfsNode *ipfsCore.IpfsNode, db *Sql, token string, clh vo.ChatListenHandler) error {

	sugar.Log.Info("Enter ChatListenMsgBlocked Function")

	if token == "" {
		listenUserId = ""
		sugar.Log.Info("Anonymous User Listen")
	} else {
		userId, err := parseToken(token)
		if err != nil {
			sugar.Log.Errorf("parseToken failed, token: %s, error: %s \n", token, err.Error())
			return errors.New("token is invaild")
		}

		listenUserId = userId
		sugar.Log.Infof("Named User Listen %s", listenUserId)
	}

	err := updateIpfsTopicSubs(ipfsNode, getRecvTopic(listenUserId))
	if err != nil {
		sugar.Log.Error("subscribe failed")
		return fmt.Errorf("subscribe failed")
	}

	for {
		if curSub == nil {
			sugar.Log.Error("*pubsub.Subscription is nil, i will return")
			return fmt.Errorf("sub is nil")
		}
		if db == nil {
			sugar.Log.Error("*Sql is nil, i will return")
			return fmt.Errorf("sql is nil")
		}

		ctx, cancel := context.WithCancel(context.Background())
		listenCancelFunc = cancel

		sugar.Log.Debugf("subscribe current topic: %s", curSub.Topic())
		data, err := curSub.Next(ctx)
		if err != nil {
			sugar.Log.Error("sub.Next failed:", err)
			time.Sleep(200 * time.Millisecond)
			continue
		}

		if data == nil {
			sugar.Log.Error("*pubsub.Message is nil, i will return")
			return fmt.Errorf("sub2 is nil")
		}

		if listenUserId == "" {
			sugar.Log.Info("listenUserId empty continue")
			continue
		}

		msg := vo.ChatPacketParams{}

		err = json.Unmarshal(data.Data, &msg)
		if err != nil {
			sugar.Log.Error("json.Unmarshal failed:", err)
			continue
		}

		if msg.Type == vo.MSG_TYPE_RECORD {

			var tmp vo.ChatSwapRecordParams
			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed:", err)
				continue
			}

			if tmp.ToId != listenUserId { // not me
				continue
			}

			sugar.Log.Debugf("record receive: %s\n", data.Data)

			res, err := handleAddRecordMsg(db, tmp)
			if err != nil {
				if err != vo.ErrorRowIsExists {
					sugar.Log.Error("handle add record failed.", err)
				}
				continue
			}

			msg.Data = res
			jsonStr, err := json.Marshal(msg)
			if err != nil {
				sugar.Log.Error("json.Marshal failed.", err)
				continue
			}
			clh.HandlerChat(string(jsonStr))

		} else if msg.Type == vo.MSG_TYPE_NEW {

			var tmp vo.ChatSwapMsgParams
			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed:", err)
				continue
			}

			if tmp.ToId != listenUserId { // not me
				continue
			}

			sugar.Log.Debugf("message receive: %s\n", data.Data)

			res, err := handleNewMsg(db, tmp)
			if err != nil {
				if err != vo.ErrorRowIsExists {
					sugar.Log.Error("handle add message failed.", err)
				} else {
					sugar.Log.Info("handle add message failed.", err)
				}
				continue
			}
			msg.Data = res
			jsonStr, err := json.Marshal(msg)
			if err != nil {
				sugar.Log.Error("json.Marshal failed.", err)
				continue
			}

			ackMsg := vo.ChatSwapAckParams{
				Type:   msg.Type,
				Id:     tmp.Id,
				FromId: tmp.ToId,
				ToId:   tmp.FromId,
			}

			err = sendMsgAck(ipfsNode, db, ackMsg)
			if err != nil {
				sugar.Log.Error("sendMsgAck failed.", err)
				// 只记录日志，继续允许
			}

			clh.HandlerChat(string(jsonStr))

		} else if msg.Type == vo.MSG_TYPE_WITHDRAW {

			var tmp vo.ChatSwapWithdrawMsgParams

			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed:", err)
				continue
			}

			if tmp.ToId != listenUserId {
				// not me
				continue
			}

			sugar.Log.Debugf("message receive: %s\n", data.Data)

			res, err := handleWithdrawMsg(db, tmp)
			if err != nil {
				sugar.Log.Error("handle withdraw message failed.", err)
				continue
			}
			msg.Data = res
			jsonStr, err := json.Marshal(msg)
			if err != nil {
				sugar.Log.Error("json.Marshal failed:", err)
				continue
			}

			clh.HandlerChat(string(jsonStr))

		} else if msg.Type == vo.MSG_TYPE_ACK {
			var tmp vo.ChatSwapAckParams

			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed:", err)
				continue
			}

			if tmp.ToId != listenUserId {
				// not me
				continue
			}

			sugar.Log.Debugf("message receive: %s\n", data.Data)

			err = handleAck(db, tmp)
			if err != nil {
				sugar.Log.Error("handle ack failed.", err)
				continue
			}

			jsonStr, err := json.Marshal(msg)
			if err != nil {
				sugar.Log.Error("json.Marshal failed.", err)
				continue
			}

			// 消息回调
			clh.HandlerChat(string(jsonStr))

		} else if msg.Type == vo.MSG_TYPE_HEARTBEAT {
			var tmp vo.FriendSwapOnlineParams

			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed:", err)
				continue
			}

			if tmp.ToId != listenUserId {
				// not me
				continue
			}

			sugar.Log.Debugf("message receive: %s\n", data.Data)

			err = handleHeartbeat(db, tmp)
			if err != nil {
				sugar.Log.Error("handle ack failed.", err)
				continue
			}

			jsonStr, err := json.Marshal(msg)
			if err != nil {
				sugar.Log.Error("json.Marshal failed.", err)
				continue
			}

			// 消息回调
			clh.HandlerChat(string(jsonStr))

		} else {
			sugar.Log.Error("unsupport msg type: ", msg.Type)
			continue
		}
	}
}

func parseToken(token string) (string, error) {
	claim, b := jwt.JwtVeriyToken(token)
	if !b {
		return "", errors.New("token is invaild. ")
	}
	userId, ok := claim["id"]
	if !ok {
		return "", fmt.Errorf("can not get userid from token: %s", token)
	}
	userIdStr, ok := userId.(string)
	if !ok {
		return "", fmt.Errorf("not string type userid: %v", userId)
	}

	return userIdStr, nil
}

func updateIpfsTopicSubs(ipfsNode *ipfsCore.IpfsNode, topic string) error {

	var err error

	if ipfsNode == nil {
		return errors.New("ipfsCore.IpfsNode is nil")
	}

	ipfsTopic, ok := TopicJoin.Load(topic)
	if !ok || ipfsTopic == nil {
		ipfsTopic, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join failed:", err)
			return fmt.Errorf("PubSub.Join failed: %s", err.Error())
		}

		TopicJoin.Store(topic, ipfsTopic)
	}

	sub, err := ipfsTopic.Subscribe()
	if err != nil {
		sugar.Log.Error("subscribe failed")
		return fmt.Errorf("subscribe failed")
	}

	// sugar.Log.Debugf("Current Sub Topic: %s", topic)

	// 更新 subscription
	curSub = sub

	return nil
}

func sendMsgAck(ipfsNode *ipfsCore.IpfsNode, db *Sql, ackMsg vo.ChatSwapAckParams) error {

	var err error

	msgTopicKey := getRecvTopic(ackMsg.ToId)

	ipfsTopic, ok := TopicJoin.Load(msgTopicKey)
	if !ok {
		ipfsTopic, err = ipfsNode.PubSub.Join(msgTopicKey)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}

		TopicJoin.Store(msgTopicKey, ipfsTopic)
	}

	msg := vo.ChatPacketParams{
		Type: vo.MSG_TYPE_ACK,
		From: ipfsNode.Identity.String(),
		Data: ackMsg,
	}
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		sugar.Log.Error("marshal send msg failed.", err)
		return err
	}

	err = ipfsTopic.Publish(context.Background(), msgBytes)
	if err != nil {
		sugar.Log.Error("ChatSendMsg failed.", err)
		return err
	}

	sugar.Log.Debugf("sendMsgAck topic: %s, data: %v", msgTopicKey, msg)

	return nil
}

func handleAck(db *Sql, msg vo.ChatSwapAckParams) error {

	if msg.Type != vo.MSG_TYPE_NEW {
		return fmt.Errorf("unsupport msg ack type: %s", msg.Type)
	}

	// msg.FromId -> 消息接收者， msg.ToId -> 消息发送者
	res, err := db.DB.Exec("update chat_msg set send_state = 1 where id = ? and from_id = ? and to_id = ?", msg.Id, msg.ToId, msg.FromId)
	if err != nil {
		sugar.Log.Error("update chat_msg ack fail, err:", err)
		return err
	}

	ar, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("update chat_msg ack fail2, err:", err)
		return err
	}

	if ar <= 0 {
		sugar.Log.Error("update chat_msg ack fail3, err: affected rows <= 0")
		return err
	}

	return nil
}

func handleHeartbeat(db *Sql, msg vo.FriendSwapOnlineParams) error {
	return nil
}

// handleAddRecordMsg 创建会话
func handleAddRecordMsg(db *Sql, msg vo.ChatSwapRecordParams) (vo.ChatRecordInfo, error) {

	defer func() {
		if r := recover(); r != nil {
			sugar.Log.Error("handleAddRecordMsg panic occurent, err:", r)
		}
	}()

	ret := vo.ChatRecordInfo{
		Id:      msg.Id,
		Name:    msg.Name,
		Img:     msg.Img,
		FromId:  msg.FromId,
		Toid:    msg.ToId,
		Ptime:   msg.Ptime,
		LastMsg: msg.LastMsg,

		UserName: "",
		Phone:    "",
		PeerId:   "",
		NickName: "",
		Sex:      0,
	}

	var ptime int64

	err := db.DB.QueryRow("SELECT ptime FROM chat_record WHERE id = ?", ret.Id).Scan(&ptime)

	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?)",
			ret.Id, "", ret.FromId, ret.Toid, ret.Ptime, ret.LastMsg)
		if err != nil {
			return ret, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return ret, err
		}

		// 查询对方信息
		err = db.DB.QueryRow("SELECT peer_id, name, phone, sex, nickname, img FROM sys_user WHERE id = ?", msg.FromId).Scan(&ret.PeerId, &ret.UserName, &ret.Phone, &ret.Sex, &ret.NickName, &ret.Img)
		if err != nil {
			if err == bsql.ErrNoRows {
				sugar.Log.Warn("not found peer info, so set empty")
			} else {
				sugar.Log.Error("query peer info failed.Err is ", err)
				return ret, err
			}
		}

		return ret, nil
	case nil:
		if ptime > msg.Ptime {
			res, err := db.DB.Exec("UPDATE chat_record SET from_id = ?, to_id = ?, ptime = ? WHERE id = ?", ret.FromId, ret.Toid, msg.Ptime, ret.Id)
			if err != nil {
				return ret, err
			}
			num, err := res.RowsAffected()
			if err != nil {
				return ret, err
			} else if num == 0 {
				return ret, err
			}
		}
		return ret, vo.ErrorRowIsExists

	default:
		return ret, err
	}

}

// handleWithdrawMsg 撤销消息
func handleWithdrawMsg(db *Sql, msg vo.ChatSwapWithdrawMsgParams) (ChatMsg, error) {

	defer func() {
		if r := recover(); r != nil {
			sugar.Log.Error("handleWithdrawMsg panic occurent, err:", r)
		}
	}()

	ret := ChatMsg{
		Id:     msg.MsgId,
		FromId: msg.FromId,
		ToId:   msg.ToId,
	}

	err := db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", ret.Id).Scan(&ret.Id, &ret.ContentType, &ret.Content, &ret.FromId, &ret.ToId, &ret.Ptime, &ret.IsWithdraw, &ret.IsRead, &ret.RecordId)

	switch err {
	case bsql.ErrNoRows:
		return ret, err
	case nil:
		if ret.IsWithdraw == 1 {
			// 已经处理
			return ret, vo.ErrorRepeatHandle
		}
		res, err := db.DB.Exec("UPDATE chat_msg SET is_with_draw = 1 WHERE id = ? and from_id = ?", ret.Id, ret.FromId)
		if err != nil {
			return ret, err
		}
		num, err := res.RowsAffected()
		if err != nil {
			return ret, err
		} else if num == 0 {
			return ret, vo.ErrorAffectZero
		}

		msgStr := "撤回了一条消息"

		res, err = db.DB.Exec("UPDATE chat_record SET last_msg = ? WHERE id = ?", msgStr, ret.RecordId)
		if err != nil {
			return ret, err
		}
		_, err = res.RowsAffected()
		if err != nil {
			return ret, err
		}

		ret.IsWithdraw = 1
		return ret, nil
	default:
		return ret, err
	}
}

// handleNewMsg 新增消息
func handleNewMsg(db *Sql, msg vo.ChatSwapMsgParams) (ChatMsg, error) {

	defer func() {
		if r := recover(); r != nil {
			sugar.Log.Error("handleNewMsg panic occurent, err:", r)
		}
	}()

	var recordId string

	ret := ChatMsg{
		Id:          msg.Id,
		ContentType: msg.ContentType,
		Content:     msg.Content,
		FromId:      msg.FromId,
		ToId:        msg.ToId,
		Ptime:       time.Now().Unix(),
		IsWithdraw:  msg.IsWithdraw,
		IsRead:      msg.IsRead,
		RecordId:    msg.RecordId,
	}

	// 检查房间是否存在
	err := db.DB.QueryRow("SELECT id FROM chat_record WHERE id = ?", ret.RecordId).Scan(&recordId)
	switch err {
	case bsql.ErrNoRows:
		ftid := strings.Split(ret.RecordId, "_")
		if len(ftid) < 2 {
			return ret, errors.New("recorId error: " + ret.RecordId)
		}

		res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) values (?, ?, ?, ?, ?, ?)",
			ret.RecordId, "", ftid[0], ftid[1], ret.Ptime, ret.Content)
		if err != nil {
			return ret, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return ret, err
		}
	case nil:
		// nothing
	default:
		return ret, err
	}

	// 检查消息是否重复
	err = db.DB.QueryRow("SELECT id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id FROM chat_msg WHERE id = ?", ret.Id).Scan(&ret.Id, &ret.ContentType, &ret.Content, &ret.FromId, &ret.ToId, &ret.Ptime, &ret.IsWithdraw, &ret.IsRead, &ret.RecordId)
	switch err {
	case bsql.ErrNoRows:
		res, err := db.DB.Exec("INSERT INTO chat_msg (id, content_type, content, from_id, to_id, ptime, is_with_draw, is_read, record_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
			ret.Id, ret.ContentType, ret.Content, ret.FromId, ret.ToId, ret.Ptime, ret.IsWithdraw, ret.IsRead, ret.RecordId)
		if err != nil {
			return ret, err
		}
		_, err = res.LastInsertId()
		if err != nil {
			return ret, err
		}

		_, err = db.DB.Exec("UPDATE chat_record SET last_msg = ?, ptime = ? WHERE id = ?", ret.Content, ret.Ptime, ret.RecordId)
		if err != nil {
			return ret, err
		}

		return ret, nil

	case nil:
		return ret, vo.ErrorRowIsExists
	default:
		return ret, err
	}
}
