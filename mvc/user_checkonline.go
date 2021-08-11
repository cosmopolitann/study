package mvc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	"github.com/cosmopolitann/clouddb/jwt"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

func FriendCheckOnline(ipfsNode *ipfsCore.IpfsNode, value string) error {

	// 用户在线检查参数
	var msg vo.FriendCheckOnlineParams

	sugar.Log.Debug("Request Param:", value)

	err := json.Unmarshal([]byte(value), &msg)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", msg)

	//校验 token 是否 满足
	claim, b := jwt.JwtVeriyToken(msg.Token)
	if !b {
		return errors.New("token 失效")
	}
	sugar.Log.Info("claim := ", claim)
	fromUserId := claim["id"].(string)

	for _, toUserId := range msg.FriendIds {
		if fromUserId == toUserId {
			continue
		}

		msgPacket := vo.ChatPacketParams{
			Type: vo.MSG_TYPE_HEARTBEAT,
			Data: vo.FriendSwapOnlineParams{
				FromId: fromUserId,
				ToId:   toUserId,
			},
			From: ipfsNode.Identity.Pretty(),
		}

		msgTopicKey := getRecvTopic(toUserId)

		ipfsTopic, err := GetPubsubTopic(ipfsNode, msgTopicKey)
		if err != nil {
			sugar.Log.Error("GetPubsubTopic failed.", err)
			return err
		}

		// // ----------------------
		// ipfsTopic, ok := TopicJoin.Load(msgTopicKey)
		// if !ok {
		// 	ipfsTopic, err = ipfsNode.PubSub.Join(msgTopicKey)
		// 	if err != nil {
		// 		sugar.Log.Error("PubSub.Join .Err is", err)
		// 		return err
		// 	}

		// 	TopicJoin.Store(msgTopicKey, ipfsTopic)
		// }
		// // ----------------------

		msgBytes, err := json.Marshal(msgPacket)
		if err != nil {
			sugar.Log.Error("marshal ChatPacketParams failed.", err)
			return err
		}

		err = ipfsTopic.Publish(context.Background(), msgBytes)
		if err != nil {
			sugar.Log.Error("CheckUserOnline Publish failed.", err)
			return err
		}

		sugar.Log.Debugf("ChatSendMsg topic: %s, userid: %s, data: %v", msgTopicKey, toUserId, msgPacket)
	}

	return nil
}
