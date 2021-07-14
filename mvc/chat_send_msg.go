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

	msgBytes, err := json.Marshal(map[string]interface{}{
		"type": vo.MSG_TYPE_NEW,
		"from": ipfsNode.Identity.String(),
		"data": swapMsg,
	})
	if err != nil {
		sugar.Log.Error("marshal send msg failed.", err)
		return ret, err
	}

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
		sugar.Log.Error("ChatSendMsg failed.", err)
		return ret, err
	}

	sugar.Log.Info("ChatSendMsg success")

	// publishUserInfo(ipfsNode, db, userId)

	// 发布消息
	return ret, nil
}

func publishUserInfo(ipfsNode *ipfsCore.IpfsNode, db *Sql, userId string) error {
	var err error
	topic := "/db-online-sync"
	// publish msg
	sugar.Log.Info("Publish Topic: ", "/db-online-sync")
	ctx := context.Background()
	tp, ok := TopicJoin.Load(topic)
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}
		TopicJoin.Store(topic, tp)
	}

	var dl vo.RespSysUser
	rows, err := db.DB.Query("select id,IFNULL(peer_id,'null'),IFNULL(name,'null'),IFNULL(phone,'null'),IFNULL(sex,0),IFNULL(ptime,0),IFNULL(utime,0),IFNULL(nickname,'null'),IFNULL(img,'null') from sys_user where id=?", userId)
	if err != nil {
		sugar.Log.Error("AddUser Query data is failed.Err is ", err)
		return err
	}
	// 释放锁
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.PeerId, &dl.Name, &dl.Phone, &dl.Sex, &dl.Ptime, &dl.Utime, &dl.NickName, &dl.Img)
		if err != nil {
			sugar.Log.Error("AddUser Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info(" AddUser Query a entire data is ", dl)
	}
	//the first step.
	var s3 UserAd
	s3.Type = "receiveUserRegister"
	s3.Data = dl
	s3.FromId = ipfsNode.Identity.String()
	//marshal UserAd.
	//the second step
	sugar.Log.Info("--- second step ---")

	jsonBytes, err := json.Marshal(s3)
	if err != nil {
		sugar.Log.Error("Publish msg is failed.Err:", err)
		return err
	}
	sugar.Log.Info("Frwarding information:=", string(jsonBytes))
	sugar.Log.Info("Local PeerId :=", ipfsNode.Identity.String())
	//the  third  step .
	sugar.Log.Info("--- third step ---")

	err = tp.Publish(ctx, jsonBytes)
	if err != nil {
		sugar.Log.Error("Publish Err:", err)
		return err
	}

	return nil
}
