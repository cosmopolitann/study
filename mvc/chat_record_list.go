package mvc

import (
	"database/sql"
	"encoding/json"

	"github.com/cosmopolitann/clouddb/jwt"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
)

func ChatRecordList(db *Sql, value string) ([]vo.ChatRecordRespListParams, error) {

	var req vo.ChatRecordListParams
	var ret []vo.ChatRecordRespListParams

	sugar.Log.Debug("Request Param: ", value)
	err := json.Unmarshal([]byte(value), &req)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return ret, err
	}
	sugar.Log.Info("Marshal data is  ", req)

	//token verify
	claim, b := jwt.JwtVeriyToken(req.Token)
	if !b {
		return ret, err
	}

	sugar.Log.Info("claim := ", claim)

	userId := claim["UserId"].(string)

	var user SysUser

	err = db.DB.QueryRow("SELECT id, peer_id, name, nickname, phone, sex, img FROM sys_user WHERE id = ?", userId).Scan(&user.Id, &user.PeerId, &user.Name, &user.NickName, &user.Phone, &user.Sex, &user.Img)
	if err != nil {
		sugar.Log.Error("query user info failed.Err is ", err)
		return ret, err
	}

	sugar.Log.Debugf("Get User: %#v", user)

	// 查询会话列表
	rows, err := db.DB.Query("SELECT id, from_id, to_id, ptime, last_msg FROM chat_record WHERE from_id = ? OR to_id = ?", user.Id, user.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
	}
	// 释放锁
	defer rows.Close()

	for rows.Next() {
		var ri vo.ChatRecordRespListParams
		err := rows.Scan(&ri.Id, &ri.FromId, &ri.ToId, &ri.Ptime, &ri.LastMsg)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return ret, err
		}

		sugar.Log.Debug(ri)

		peerId := ""

		if ri.FromId == user.Id {
			peerId = ri.ToId

			ri.FromName = user.Name
			ri.FromImg = user.Img
			ri.FromPhone = user.Phone
			ri.FromPeerId = user.PeerId
			ri.FromNickName = user.NickName
			ri.FromSex = user.Sex

		}
		if ri.ToId == user.Id {
			peerId = ri.FromId

			ri.ToName = user.Name
			ri.ToImg = user.Img
			ri.ToPhone = user.Phone
			ri.ToPeerId = user.PeerId
			ri.ToNickName = user.NickName
			ri.ToSex = user.Sex

		}

		sugar.Log.Debugf("Get Record %#v", ri)

		if len(peerId) != 0 {
			var peer SysUser
			err = db.DB.QueryRow("SELECT id, peer_id, name, nickname, phone, sex, img FROM sys_user WHERE id = ?", peerId).Scan(&peer.Id, &peer.PeerId, &peer.Name, &peer.NickName, &peer.Phone, &peer.Sex, &peer.Img)
			if err != nil {
				if err == sql.ErrNoRows {
					sugar.Log.Warn("not found peer info, so set empty")
				} else {
					sugar.Log.Error("query peer info failed.Err is ", err)
					return ret, err
				}

			} else {
				sugar.Log.Debugf("Update Peer: %#v", peer)

				ri.Name = peer.Name
				ri.Img = peer.Img

				if ri.FromId == peer.Id {
					ri.FromName = peer.Name
					ri.FromImg = peer.Img
					ri.FromPhone = peer.Phone
					ri.FromPeerId = peer.PeerId
					ri.FromNickName = peer.NickName
					ri.FromSex = peer.Sex
				} else {
					ri.ToName = peer.Name
					ri.ToImg = peer.Img
					ri.ToPhone = peer.Phone
					ri.ToPeerId = peer.PeerId
					ri.ToNickName = peer.NickName
					ri.ToSex = peer.Sex
				}
			}
		}

		err = db.DB.QueryRow("SELECT count(id) FROM chat_msg WHERE record_id = ? AND is_read = 0 AND from_id != ?", ri.Id, user.Id).Scan(&ri.UnreadMsgNum)
		if err != nil {
			sugar.Log.Warn("Query unread msg failed.Err is ", err)
		}

		ret = append(ret, ri)
	}

	return ret, nil

}
