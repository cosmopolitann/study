package mvc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/mr-tron/base58/base58"
	"log"
	"strconv"
	"time"
)

func SyncUser(db *Sql, value string) error {

	//var user SysUser
	//err := json.Unmarshal([]byte(value), &user)
	//if err != nil {
	//	sugar.Log.Error("---同步 解析 数据 失败 ---:", err)
	//	return err
	//}
	//sugar.Log.Info("params ：= ", user)
	//
	//
	//t := time.Now().Unix()
	//stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?)")
	//if err != nil {
	//	sugar.Log.Error("同步 Insert data to sys_user is failed.")
	//	return err
	//}
	//
	////sid := strconv.FormatInt(user.Id, 10)
	//res, err := stmt.Exec(user.Id, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName)
	//if err != nil {
	//	sugar.Log.Error("同步 Insert data to sys_user is failed.", res)
	//	return err
	//}
	//c, _ := res.RowsAffected()
	//sugar.Log.Info("~~~~~  同步   into sys_user data is Successful ~~~~~~", c)
	////生成 token
	//// 手机号
	////token,err:=jwt.GenerateToken(user.Phone,60)
	//
	//return nil
	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {

	}
	sugar.Log.Info("params ：= ", user)

	l, e := FindIsExistUser(db, user)
	if e != nil {
		sugar.Log.Error("FindIsExistUser info is Failed.")
	}
	// l > 0 user is exist.
	sugar.Log.Error("-----------1")

	if l > 0 {
		sugar.Log.Error("user is exist.")
		return errors.New("user is exist.")
	}

	//inExist insert into sys_user.

	sugar.Log.Info("-----------用户 信息 ", user)

	id := utils.SnowId()
	//create now time
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.")
		return err
	}
	sid := strconv.FormatInt(id, 10)
	res, err := stmt.Exec(sid, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.", res)
		return err
	}
	c, _ := res.RowsAffected()
	sugar.Log.Info("~~~~~   Insert into sys_user data is Successful ~~~~~~", c)
	return nil
}

// 文章

func SyncArticle(db *Sql, value string) error {
	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New("解析字段错误")
	}
	sugar.Log.Info("Marshal data is  ", art)
	id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return errors.New("插入article 表数据 失败")
	}
	sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(sid, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return errors.New("插入数据失败")
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New("插入数据失败")
	}
	return nil
}

// 同步 文章 播放量

func SyncAticlePlay(db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}
	vl, _ := rows.Columns()
	sugar.Log.Info("vl ", vl)

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set play_num=? where id=?")
	if err != nil {
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.PlayNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	return nil
}
func SyncArticleShareAdd(db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("同步 Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("同步 Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("同步 Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("同步 Query data is failed.Err is ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("同步 Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("同步 Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error("同步 Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("同步 Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("同步 Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("同步 Update  is failed.The err is ", err)
		return err
	}

	return nil
}

func SyncUserRegister(db *Sql, value string) error {
	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("同步 Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("同步 Marshal data is  ", art)
	if err != nil {
		sugar.Log.Error("同步 Insert into article table is failed.", err)
		return err
	}
	//select the data is exist.
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("同步 Query data is failed.Err is ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("同步 Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("同步 Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" 同步 update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error("同步 Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("同步 Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("同步 Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("同步 Update  is failed.The err is ", err)
		return err
	}

	return nil

}
// // 同步 文章 播放量

func SyncArticleShare(db *Sql, value string) error {

	var dl Article
	var art vo.ArticlePlayAddParams
	err := json.Unmarshal([]byte(value), &art)

	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
	}
	sugar.Log.Info("Marshal data is  ", art)

	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	return nil
}

func SyncUserUpdate(db *Sql, value string) error {

	return nil
}


var Topicmp map[string]*pubsub.Topic

func SyncTopicData(ipfsNode *ipfsCore.IpfsNode, db *Sql, value string) error {
	//监听topic
	topic := "/db-online-sync"
	sugar.Log.Info("开始监听主题 : ", topic)
	sugar.Log.Info("subscrib topic: ", topic)

	ctx := context.Background()
	sugar.Log.Info("加入 主题 房间  : ", topic)
	// 判断 map 是否存在 当前 主题

	tp, err := ipfsNode.PubSub.Join(topic)
	if err != nil {
		sugar.Log.Error("subscribe Join failed.", err)
		return err
	}
	//
	sugar.Log.Info("将tp 加入 到 map中  : ", topic)
	Topicmp = make(map[string]*pubsub.Topic)
	Topicmp["/db-online-sync"] = tp
	sugar.Log.Info("主题map :", Topicmp)

	sugar.Log.Info(" Subscribe topic  tp :", tp)

	sub, err := tp.Subscribe()
	if err != nil {
		sugar.Log.Error("subscribe failed.", err)
		return err
	}

	for {
		sugar.Log.Info("------------------------------------------------")
		sugar.Log.Info("开始 同步 消息")

		data, err := sub.Next(ctx)
		if err != nil {
			sugar.Log.Error("subscribe failed.", err)
			continue
		}
		msg := data.Message
		log.Println("------ 收到的消息的内容---",msg.Data)

		log.Printf("------ 收到的消息的类型 %T\n----",msg.Data)
		fromId := msg.From
		sugar.Log.Info("-----来自谁的消息-----:", string(fromId))

		peerId := ipfsNode.Identity.String()

		sugar.Log.Info("本地节点peerId:", peerId)

		//
		var recieve vo.SyncMsgParams
		err = json.Unmarshal(msg.Data, &recieve)
		if err != nil {
			sugar.Log.Error("解析失败:", err)
			continue
		}
		sugar.Log.Info("---- 解析收到同步消息是---:", recieve)
		sugar.Log.Info("---- 这是 谁发来的消息 fromId    peerId:", string(fromId))
		sugar.Log.Info("---- 这是消息内容里面的 recieve.FromId    ---:", recieve.FromId)
		sugar.Log.Info("---- 本机  peerId ---:", peerId)
		sugar.Log.Infof("----这是  string(fromId)  %v \n---:", string(fromId))
		sugar.Log.Info("---- 这是 谁发来的消息 fromId    msg.From:", msg.From)
		sugar.Log.Infof("----这是  string(fromId)的类型  %T \n---:", string(fromId))
		log.Printf("公共节点的信息 \n",msg.From)



		wayId:="12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J"

		sugar.Log.Info("----公共网关节点 id =---:", wayId)

		FromID:=base58.Encode(fromId)

		log.Println("-----这是 base58 转string 类型 ----FromID ;",FromID)

		if FromID==wayId{
				sugar.Log.Info("---- 因为 公共网关 节点id 等于 i8j 所以满足条件进来 ---:", peerId)

				
				if peerId==recieve.FromId{

					sugar.Log.Info("---- 因为 本地 节点id 等于 recieve fromId  所以不满足 ---:")
					sugar.Log.Info("---- 因为 本地 节点id  ---",peerId)
					sugar.Log.Info("---- 因为 本地 recieve.FromId  ---",recieve.FromId)

					continue
				}else {

					if recieve.Method == "receiveArticleAdd" {
						//  添加 文章  入库
						//第一步 解析

						var syn vo.SyncRecieveArticleParams
						err = json.Unmarshal(msg.Data, &syn)
						if err != nil {
							sugar.Log.Error("同步 解析 用户字段 错误:", err)
							continue
						}
						// string

						userInfo, err := json.Marshal(syn.Data)
						if err != nil {
							sugar.Log.Error("同步添加文章失败:", err)
							continue
						}
						sugar.Log.Info("解析收到 同步消息的receiveArticleAdd 消息是", recieve.Method)
						err = db.SyncArticle(string(userInfo))
						if err != nil {
							sugar.Log.Error("同步添加文章失败:", err)
							continue
						}

						sugar.Log.Info("同步添加文章成功")
					} else if recieve.Method == "receiveArticlePlayAdd" {
						sugar.Log.Info("-----  同步增加播放次数  -----")

						sugar.Log.Info("-----  同步增加播放次数 的数据  -----", value)
						//---
						//第一步 解析

						var syn vo.SyncRecievePlayParams
						err = json.Unmarshal(msg.Data, &syn)
						if err != nil {
							sugar.Log.Error("同步 解析 用户字段 错误:", err)
							continue
						}
						// string
						userInfo, err := json.Marshal(syn.Data)
						if err != nil {
							sugar.Log.Error("同步 播放 数量 失败:", err)
							continue
						}
						sugar.Log.Info("解析收到 receiveArticlePlayAdd 消息类型是", recieve.Method)

						sugar.Log.Info("解析收到 receiveArticlePlayAdd 消息内容是", string(userInfo))

						err = db.SyncArticlePlay(string(userInfo))
						if err != nil {
							sugar.Log.Error("-----  同步增加播放次数 失败  -----", err)
							continue
						}
						//  增加播放次数
						//var tmp vo.ChatMsgParams
						//json1, _ := json.Marshal(msg.Data)
						//json.Unmarshal(json1, &tmp)
						//
						//res, err := handleNewMsg(db, tmp)
						//if err != nil {
						//	sugar.Log.Error("handle add message failed.", err)
						//	continue
						//}
						//msg.Data = res
						//jsonStr, _ := json.Marshal(msg)
						//clh.HandlerChat(string(jsonStr))
					} else if recieve.Method == "receiveArticleShareAdd" {
						//  增加 分享 次数

						sugar.Log.Info("-----  同步  增加 分享 次数  -----")

						sugar.Log.Info("-----  同步  增加 分享 次数  的数据  -----", value)
						//--
						//第一步 解析

						var syn vo.SyncRecievePlayParams
						err = json.Unmarshal(msg.Data, &syn)
						if err != nil {
							sugar.Log.Error("同步 解析 用户字段 错误:", err)
							continue
						}
						// string
						userInfo, err := json.Marshal(syn.Data)
						if err != nil {
							sugar.Log.Error("同步 播放 数量 失败:", err)
							continue
						}
						sugar.Log.Info("解析收到 receiveArticlePlayAdd 消息是", recieve.Method)

						//----
						err = db.SyncArticleShareAdd(string(userInfo))
						if err != nil {
							sugar.Log.Error("-----  同步  增加 分享 次数  失败  -----", err)
							continue
						}
						sugar.Log.Info(" 增加 分享 次数")

					} else if recieve.Method == "receiveUserRegister" {
						// 添加用户 信息
						sugar.Log.Info("-----  同步  添加用户 信息  -----")

						sugar.Log.Info("-----  同步  添加用户 信息  -----", value)

						//----

						//第一步 解析

						var syn vo.SyncRecieveUsesrParams
						err = json.Unmarshal(msg.Data, &syn)
						if err != nil {
							sugar.Log.Error("同步 解析 用户字段 错误:", err)
							continue
						}
						// string
						userInfo, err := json.Marshal(syn.Data)
						if err != nil {
							sugar.Log.Error("同步 播放 数量 失败:", err)
							continue
						}
						sugar.Log.Info("解析收到 receiveArticlePlayAdd 消息是", recieve.Method)

						//-------
						err = db.SyncUser(string(userInfo))
						if err != nil {
							sugar.Log.Error("----- 添加用户 信息 失败  -----", err)
							continue
						}
						sugar.Log.Info(" 添加用户 信息 成功")
					} else {
						sugar.Log.Info("不满足条件，继续:")
						continue
					}
				}
			}else {
				sugar.Log.Info("不满足条件，继续:")

				continue
			}

	}
	return nil
}
