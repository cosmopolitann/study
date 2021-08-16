package mvc

import (
	"database/sql"
	"encoding/json"
	"errors"
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
	rows, err := db.DB.Query("SELECT id, from_id, to_id, ptime, last_msg FROM chat_record WHERE from_id = ? OR to_id = ? ORDER BY ptime DESC", user.Id, user.Id)
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

		peer, err := apis.GetUserInfo(req.Token, peerId)
		if err != nil {
			sugar.Log.Error("query peer info failed.Err is ", err)
			return ret, err
		}

		if err != nil {

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
				peer.Nickname = fnickname
				ri.Name = fnickname
			}

			if req.Keyword != "" {
				if peer.Nickname == "" {
					continue
				} else {
					re, err := regexp.Compile(".*" + regexp.QuoteMeta(req.Keyword) + ".*")
					if err != nil {
						return ret, err
					}

					if !re.Match([]byte(peer.Nickname)) {
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
				} else if peer.Nickname != "" && re.Match([]byte(peer.Nickname)) {
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
				ri.FromNickName = peer.Nickname
				ri.FromSex = peer.Sex
			} else {
				ri.ToName = peer.Name
				ri.ToImg = peer.Img
				ri.ToPhone = peer.Phone
				ri.ToPeerId = peer.PeerId
				ri.ToNickName = peer.Nickname
				ri.ToSex = peer.Sex
			}

			if fnickname != "" {
				peer.Nickname = fnickname
				ri.Name = fnickname

			} else if peer.Nickname != "" {
				ri.Name = peer.Nickname

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
