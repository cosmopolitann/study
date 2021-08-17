package mvc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/cosmopolitann/clouddb/apis"
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

	userId := claim["id"].(string)

	if len(req.CustomerId) > 0 && req.CustomerId != userId {
		var rid string

		recordId := genRecordID(userId, req.CustomerId)
		err = db.DB.QueryRow("SELECT id FROM chat_record where id = ?", recordId).Scan(&rid)
		if err != nil && err != sql.ErrNoRows {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return ret, errors.New("查询下载列表信息失败")
		}

		if rid == "" {
			// no room
			res, err := db.DB.Exec("INSERT INTO chat_record (id, name, from_id, to_id, ptime, last_msg) VALUES (?, ?, ?, ?, ?, ?)", recordId, "", userId, req.CustomerId, time.Now().Unix(), "")
			if err != nil {
				sugar.Log.Error("INSERT INTO chat_record is Failed.", err)
				return ret, err
			}

			_, err = res.LastInsertId()
			if err != nil {
				sugar.Log.Error("INSERT INTO chat_record is Failed2.", err)
				return ret, err
			}
		}
	}

	user, err := apis.GetUserInfo(req.Token, userId)
	if err != nil {
		sugar.Log.Error("query user info failed.Err is ", err)
		return ret, err
	}

	sugar.Log.Debugf("Get User: %#v", user)

	// 查询会话列表
	rows, err := db.DB.Query("SELECT id, from_id, to_id, ptime, last_msg FROM chat_record WHERE from_id = ? OR to_id = ? ORDER BY ptime DESC", userId, userId)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
	}
	// 释放锁
	defer rows.Close()

	var records []vo.ChatRecored
	mapUserIds := make(map[string]string)
	for rows.Next() {
		var ri vo.ChatRecored
		err := rows.Scan(&ri.Id, &ri.FromId, &ri.ToId, &ri.Ptime, &ri.LastMsg)
		if err != nil {
			sugar.Log.Error("Query data is failed.Err is ", err)
			return ret, err
		}

		records = append(records, ri)

		mapUserIds[ri.FromId] = ri.FromId
		mapUserIds[ri.ToId] = ri.ToId
	}

	if len(mapUserIds) > 0 {
		var userIds []string
		for _, userId := range mapUserIds {
			userIds = append(userIds, userId)
		}
		uModels, err := apis.GetUsers(req.Token, userIds)
		if err == nil {
			for _, u := range uModels {
				_, err := db.DB.Exec("INSERT OR REPLACE INTO sys_user(id, peer_id, name, phone, sex, ptime, utime, nickname, img, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
					u.Id, u.PeerId, u.Name, u.Phone, u.Sex, u.Ptime, u.Utime, u.Nickname, u.Img, u.Role)
				if err != nil {
					return ret, fmt.Errorf("update user err: %w", err)
				}
			}
		} else {
			sugar.Log.Error("apis.GetUsers err:", err)
		}
	}

	for _, record := range records {
		ri := vo.ChatRecordRespListParams{
			Id:      record.Id,
			FromId:  record.FromId,
			ToId:    record.ToId,
			Ptime:   record.Ptime,
			LastMsg: record.LastMsg,
		}

		peerId := ""

		if ri.FromId == user.Id {
			peerId = ri.ToId

			ri.FromName = user.Name
			ri.FromImg = user.Img
			ri.FromPhone = user.Phone
			ri.FromPeerId = user.PeerId
			ri.FromNickName = user.Nickname
			ri.FromSex = user.Sex
		}

		if ri.ToId == user.Id {
			peerId = ri.FromId

			ri.ToName = user.Name
			ri.ToImg = user.Img
			ri.ToPhone = user.Phone
			ri.ToPeerId = user.PeerId
			ri.ToNickName = user.Nickname
			ri.ToSex = user.Sex
		}

		sugar.Log.Debugf("Get Record %#v", ri)

		var peer SysUser
		err = db.DB.QueryRow("SELECT id, peer_id, name, nickname, phone, sex, img FROM sys_user WHERE id = ?", peerId).Scan(&peer.Id, &peer.PeerId, &peer.Name, &peer.NickName, &peer.Phone, &peer.Sex, &peer.Img)
		if err != nil {
			if err != sql.ErrNoRows {
				sugar.Log.Error("query peer info failed.Err is ", err)
				return ret, err
			}

			// not found peer
			var fnickname string
			err := db.DB.QueryRow("SELECT friend_nickname FROM user_friend WHERE user_id = ? AND friend_id = ?", userId, peerId).Scan(&fnickname)
			if err != nil {
				if err != sql.ErrNoRows {
					sugar.Log.Error("query user_friend nickname failed.Err is ", err)
					return ret, err
				}
			}

			if fnickname != "" {
				peer.NickName = fnickname
				ri.Name = fnickname
			}

			if req.Keyword != "" {
				if peer.NickName == "" {
					continue
				} else {
					re, err := regexp.Compile(".*" + regexp.QuoteMeta(req.Keyword) + ".*")
					if err != nil {
						return ret, err
					}

					if !re.Match([]byte(peer.NickName)) {
						continue
					}
				}
			}

		} else {
			sugar.Log.Debugf("Update Peer: %#v", peer)

			var fnickname string
			err := db.DB.QueryRow("SELECT friend_nickname FROM user_friend WHERE user_id = ? AND friend_id = ?", userId, peerId).Scan(&fnickname)
			if err != nil {
				if err != sql.ErrNoRows {
					sugar.Log.Error("query user_friend nickname failed.Err is ", err)
					return ret, err
				}
			}

			ri.Img = peer.Img

			if req.Keyword != "" {
				re, err := regexp.Compile(".*" + regexp.QuoteMeta(req.Keyword) + ".*")
				if err != nil {
					return ret, err
				}

				matched := false
				if fnickname != "" && re.Match([]byte(fnickname)) {
					matched = true
				} else if peer.NickName != "" && re.Match([]byte(peer.NickName)) {
					matched = true
				} else if peer.Name != "" && re.Match([]byte(peer.Name)) {
					matched = true
				}

				if !matched {
					continue
				}
			}

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

			if fnickname != "" {
				peer.NickName = fnickname
				ri.Name = fnickname

			} else if peer.NickName != "" {
				ri.Name = peer.NickName

			} else if peer.Name != "" {
				ri.Name = peer.Name
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
