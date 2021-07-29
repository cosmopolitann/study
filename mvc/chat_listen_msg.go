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

var listenUserId string

func ChatListenMsgUpdateUser(token string) error {

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
	return nil
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

func GetIpfsTopic(ipfsNode *ipfsCore.IpfsNode, topicJoin *vo.TopicJoinMap, topic string) (*pubsub.Topic, error) {

	var err error

	if ipfsNode == nil {
		return nil, errors.New("ipfsCore.IpfsNode is nil")
	}

	if topicJoin == nil {
		return nil, errors.New("TopicJoin is nil")
	}

	ipfsTopic, ok := TopicJoin.Load(topic)
	if !ok || ipfsTopic == nil {
		ipfsTopic, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join failed:", err)
			return nil, fmt.Errorf("PubSub.Join failed: %s", err.Error())
		}

		TopicJoin.Store(topic, ipfsTopic)
	}

	return ipfsTopic, nil
}

func ChatListenMsgBlocked(ipfsNode *ipfsCore.IpfsNode, db *Sql, token string, clh vo.ChatListenHandler) error {

	sugar.Log.Info("Enter ChatListenMsgBlocked Function")

	defer func() {
		if r := recover(); r != nil {
			sugar.Log.Error("End ChatListenMsgBlocked panic occurent, err:", r)
		} else {
			sugar.Log.Error("End ChatListenMsgBlocked")
		}
	}()

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

	ipfsTopic, err := GetIpfsTopic(ipfsNode, TopicJoin, vo.CHAT_MSG_SWAP_TOPIC)
	if err != nil {
		sugar.Log.Error("GetIpfsTopic failed")
		return fmt.Errorf("GetIpfsTopic failed")
	}

	sub, err := ipfsTopic.Subscribe()
	if err != nil {
		sugar.Log.Error("subscribe failed")
		return fmt.Errorf("subscribe failed")
	}

	var msg vo.ChatListenParams

	ctx := context.Background()
	for {
		if sub == nil {
			sugar.Log.Error("*pubsub.Subscription is nil, i will return")
			return fmt.Errorf("sub is nil")
		}
		data, err := sub.Next(ctx)
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

		msg = vo.ChatListenParams{}

		err = json.Unmarshal(data.Data, &msg)
		if err != nil {
			sugar.Log.Error("json.Unmarshal failed1:", err)
			continue
		}

		if msg.Type == vo.MSG_TYPE_RECORD {

			var tmp vo.ChatSwapRecordParams
			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed1:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed2:", err)
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
				sugar.Log.Error("json.Marshal failed2.", err)
				continue
			}
			clh.HandlerChat(string(jsonStr))

		} else if msg.Type == vo.MSG_TYPE_NEW {

			var tmp vo.ChatSwapMsgParams
			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed3:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed3:", err)
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
				}
				continue
			}
			msg.Data = res
			jsonStr, err := json.Marshal(msg)
			if err != nil {
				sugar.Log.Error("json.Marshal failed4.", err)
				continue
			}

			clh.HandlerChat(string(jsonStr))

		} else if msg.Type == vo.MSG_TYPE_WITHDRAW {

			var tmp vo.ChatSwapWithdrawMsgParams

			json1, err := json.Marshal(msg.Data)
			if err != nil {
				sugar.Log.Error("json.Marshal failed5:", err)
				continue
			}

			err = json.Unmarshal(json1, &tmp)
			if err != nil {
				sugar.Log.Error("json.Unmarshal failed4:", err)
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
				sugar.Log.Error("json.Marshal failed6:", err)
				continue
			}

			clh.HandlerChat(string(jsonStr))

		} else {
			sugar.Log.Error("unsupport msg type: ", msg.Type)
			continue
		}
	}
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
