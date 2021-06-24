package mvc

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/utils"
	"github.com/cosmopolitann/clouddb/vo"
	shell "github.com/ipfs/go-ipfs-api"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/mr-tron/base58/base58"
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

	//id := utils.SnowId()
	//create now time
	//t:=time.Now().Format("2006-01-02 15:04:05")
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.")
		return err
	}
	//sid := strconv.FormatInt(user.Id, 10)
	res, err := stmt.Exec(user.Id, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img)
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
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Infof("这是恐慌信息:", err)
		}
	}()
	topic := "/db-online-sync"
	sugar.Log.Info("开始监听主题 : ", topic)
	sugar.Log.Info("subscrib topic: ", topic)

	ctx := context.Background()
	sugar.Log.Info("加入 主题 房间  : ", topic)
	// 判断 map 是否存在 当前 主题

	// tp, err := ipfsNode.PubSub.Join(topic)
	// if err != nil {
	// 	sugar.Log.Error("subscribe Join failed.", err)
	// 	return err
	// }
	// //
	// sugar.Log.Info("将tp 加入 到 map中  : ", topic)
	// Topicmp = make(map[string]*pubsub.Topic)
	// Topicmp["/db-online-sync"] = tp
	// sugar.Log.Info("主题map :", Topicmp)
	tp, ok := TopicJoin.Load(topic)
	var err error
	if !ok {
		tp, err = ipfsNode.PubSub.Join(topic)
		if err != nil {
			sugar.Log.Error("PubSub.Join .Err is", err)
			return err
		}
		TopicJoin.Store(topic, tp)
	}
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
		log.Println("------ 收到的消息的内容---", msg.Data)

		log.Printf("------ 收到的消息的类型 %T\n----", msg.Data)
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
		wayId := "12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J"
		sugar.Log.Info("----公共网关节点 id =---:", wayId)
		FromID := base58.Encode(fromId)
		sugar.Log.Info("-----来自谁的消息 转码的FromID-----:", string(FromID))

		if FromID == wayId {
			sugar.Log.Info("---- 因为 公共网关 节点id 等于 i8j 所以满足条件进来 ---:", peerId)

			if peerId == recieve.FromId {
				sugar.Log.Info("---- 因为 本地 节点id 等于 recieve fromId  所以不满足 ---:")
				continue
			} else {

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
		} else {
			sugar.Log.Info("不满足条件，继续:")

			continue
		}

	}
	return nil
}

var sh *shell.Shell

//Off Line Data.
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
func OffLineSyncData(db *Sql, path string) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Info("  捕捉恐慌 ~~~~~~~~~~~~1:", err)
		} else {
			sugar.Log.Info("   正常 ~~~~~~~~~~~~2")

		}

	}()
	sugar.Log.Info("--- Start excute offline task ---")
	var defaltPath = path + "local"
	sugar.Log.Info(" Local Path :", defaltPath)
	b := Exist(defaltPath)
	if !b {
		//create file.
		sugar.Log.Info(" Local File is exist and create local file. ")
		_, err1 := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
		if err1 != nil {
			sugar.Log.Error(" Create Local File Is Failed.Err: ", err1)
		}
	}
	sh = shell.NewShell("localhost:5001")
	sugar.Log.Info(" ---  Start resolve remote ipns data.  ---")
	// result, err := sh.Resolve("k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v")
	// RemoteIpnsAddr
	// sugar.Log.Info(" Ipns Addr: ", RemoteIpnsAddr)

	result, err := sh.Resolve("k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v")

	if err != nil {
		sugar.Log.Error(" Ipns Addr resolve is failed. Err:", err)
	}

	sugar.Log.Info("The result what ipfs cat remote cid. ", result)
	hash := result
	sugar.Log.Info("Remote ipns addr cid : ", hash)
	//read local file.
	sugar.Log.Info(" Read local file content. ")
	local, err := ioutil.ReadFile(path + "local") // just pass the file name
	if err != nil {
		sugar.Log.Error(" Read local file content is failed. ", err)
	}
	sugar.Log.Info(" Read local file content. ", string(local))
	//read remote file.
	sugar.Log.Info(" Start read remote file content to use remote cid. ")
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	remote1, Errremote := ioutil.ReadAll(read)
	if Errremote != nil {
		sugar.Log.Error("  Read remote file content is failed.Err: ", Errremote)
	}
	sugar.Log.Info("  Read remote file content: ", string(remote1))
	if strings.ToLower(string(local)) == strings.ToLower(string(remote1)) {
		sugar.Log.Info(" Remote equal Local ")
	} else {
		sugar.Log.Info(" Remote not equal Local ")
		// loop pull not equal cid
		// string split by  _
		sugar.Log.Info(" Split remote and local file user _  ")

		remoteStr1 := strings.Split(string(remote1), "_")
		localStr1 := strings.Split(string(local), "_")

		diff := difference(localStr1, remoteStr1)

		sugar.Log.Info(" This is diff array: ", diff)
		sugar.Log.Info(" This is diff array lenth: ", len(diff))

		for i := 1; i < len(diff); i++ {
			sugar.Log.Info(" --- loop  diff array ---- ", i)
			sugar.Log.Info(" --- diff array value ---- ", diff[i])

			cidPath := path + string(diff[i])
			sugar.Log.Info(" cidPath : ", cidPath)
			sugar.Log.Info(" Ipfs get cidpath  : ", cidPath)
			sugar.Log.Info(" Ipfs get cid hash :", string(diff[i]))
			err := sh.Get(string(diff[i]), cidPath)
			if err != nil {
				fmt.Println(err)
				sugar.Log.Error(" Ipfs get cid hash is failed.Err:", err)
			}
			sugar.Log.Info(" Read diff cid file by line .")
			sugar.Log.Info(" Open cidPath file : ", cidPath)
			f1, err := os.Open(cidPath)
			if err != nil {
				sugar.Log.Error(" Open cidPath is failed.Err:", err)
			}
			f1.Close()
			rd1 := bufio.NewReader(f1)
			for {
				sugar.Log.Info(" Start loop read cidPath file by line util end. ")
				line, err := rd1.ReadString('\n') // by '\n' as end sign of closure.
				if err != nil || io.EOF == err {
					sugar.Log.Info(" --- break ----")
					sugar.Log.Info(" ---ReadString is failed.Err:", err)
					break
				}
				sugar.Log.Info(" Data for each line:", line)
				sugar.Log.Info(" 读出每一行的数据信息 :", line)

				sugar.Log.Infof(" Data type for each line is %T .\n", line)
				// exec sql read cidPath file by line.
				sugar.Log.Info(" Start excute sql  by read cidpath file content. ", line)
				stmt, err := db.DB.Prepare(string(line))
				if err != nil {
					sugar.Log.Error("Insert data into table is failed.", err)
					//continue
				}
				res, err := stmt.Exec()
				sugar.Log.Info(" --- 开始插入数据 ---  ", string(line))
				time.Sleep(time.Second)
				if err != nil {
					sugar.Log.Error("Insert data into  is Failed.", err)
					//continue
				}
				l, err := res.RowsAffected()
				if l == 0 {
					sugar.Log.Error("Excute sql is failed.Err:", err)
					//continue
				}
			}
			// 	delete cidPath file.
			sugar.Log.Info(" Start delete cidPath file.")
			existed := true
			if _, err := os.Stat(cidPath); os.IsNotExist(err) {
				existed = false
			}
			if existed {
				err := os.Remove(cidPath)
				sugar.Log.Info(" delete cidPath file:", cidPath)

				if err != nil {
					sugar.Log.Error(" delete cidPath file is failed.Err:", err)
					sugar.Log.Error(" Delete cidPath file is failed.Err:", err)
				} else {
					sugar.Log.Info(" Delete cidPath file is successful !!! ", cidPath)
				}
			}
		}
		// delete remote file and read remote file content cid
		// write the content to this local file.
		sugar.Log.Info(" Open file defaltPath : ", defaltPath)
		local_f, err1 := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //open file.
		if err1 != nil {
			sugar.Log.Error(" Open file is failed.Err:", err1)
		}
		_, err = local_f.WriteString(string(remote1))
		if err != nil {
			sugar.Log.Error(" Write remote content to this local file is failed.Err: ", err)
		}
		sugar.Log.Info(" Write remote content to this local file is successful!! ")
		sugar.Log.Info(" Start delete file ")
		sugar.Log.Info(" Delete file path:", defaltPath+"remote")
		existed := true
		if _, err := os.Stat(defaltPath + "remote"); os.IsNotExist(err) {
			existed = false
		}
		if existed {
			sugar.Log.Info(" Delete file path:", defaltPath+"remote")
			err := os.Remove(defaltPath + "remote")
			if err != nil {
				sugar.Log.Error(" Delete file path is failed.Err:", err)
			} else {
				sugar.Log.Info(" Delete file path is successful! ", defaltPath+"remote")
			}
		}
		sugar.Log.Info("-----------Execute sql is successful. ---------------")
		// f, err := os.Open("./output")
		// if err != nil {
		// 	panic(err)
		// }
		// defer f.Close()

		// rd := bufio.NewReader(f)
		// for {
		// 	line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		// 	if err != nil || io.EOF == err {
		// 		break
		// 	}
		// 	fmt.Println(line)
		// 	fmt.Printf("类型 是 %T\n ", line)
		// 	// 执行 sql 语句 试试
		// 	fmt.Println("----- 开始 执行  sql 语句 -----")
		// 	stmt, err := db.DB.Prepare(string(line))
		// 	if err != nil {
		// 		sugar.Log.Error("Insert into cloud_file table is failed.", err)
		// 		continue
		// 	}
		// 	res, err := stmt.Exec()
		// 	if err != nil {
		// 		sugar.Log.Error("Insert into file  is Failed.", err)
		// 		continue
		// 	}
		// 	l, _ := res.RowsAffected()
		// 	if l == 0 {
		// 		sugar.Log.Error("执行sql 失败 原因:", err)
		// 		continue
		// 	}
		// }

		// // 完成之后 删除 output 文件
		// // 新建一个 cid 文件 拼接字符串
		// if checkFileIsExist("./version") { //如果文件存在
		// 	f, err1 := os.OpenFile("./version", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666) //打开文件
		// 	if err1 != nil {
		// 		fmt.Println("err", err1)

		// 	}
		// 	fmt.Println("文件存在")
		// 	//读文件 写文件信息
		// 	_, err = f.WriteString("writeString : " + "_"+)
		// 	if err != nil {
		// 		log.Println(err)
		// 		return
		// 	}

		// }
		// } else {
		// 	f, err1 := os.Create("./version") //创建文件
		// 	fmt.Println("文件不存在")
		// 	if err1 != nil {
		// 		fmt.Println("err", err1)
	}
	//ipns
	sugar.Log.Info(" Start upload cid to gateway.io ipns. ")
	UploadFile(path, hash)
	sugar.Log.Info(" Because local  =====   remote.  ")
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func difference(slice1 []string, slice2 []string) []string {
	var diff []string
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}

// local update file

func UploadFile(path string, hash string) {
	// resolve k5 => /ipfs/cid , then pull the remote file.
	sugar.Log.Info(" Start resolve k5 => /ipfs/cid  ")
	sugar.Log.Info(" Exist update file state .")

	var updateCid string
	b := Exist(path + "update")
	if !b {
		//create the update file.
		sugar.Log.Info(" Update file is exist,so create it and open at the same time.")
		_, err1 := os.OpenFile(path+"update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
		if err1 != nil {
			sugar.Log.Error(" Create update is failed .Err:", err1)
		}
	} else {
		//read the update file.
		sugar.Log.Info(" Read update file to get the content. ")
		bytes1, err := ioutil.ReadFile(path + "update")

		if err != nil {
			sugar.Log.Error(" Read update file to get the content is failed.Err:", err)
		}
		// upload the file to ipfs.
		hash_local, err := sh.Add(bytes.NewBufferString(string(bytes1)))
		if err != nil {
			sugar.Log.Error(" Upload the file to ipfs is failed.Err:", err)
		}
		sugar.Log.Info(" THe hash value what upload file to create a hash by ipfs. ", hash_local)
		updateCid = hash_local
	}

	//read remote file.
	sugar.Log.Infof(" Cat remote %s to get content by ipfs. \n", hash)
	read, err := sh.Cat(hash)
	if err != nil {
		sugar.Log.Error(" Cat remote hash is failed.Err:", err)
	}
	remote1, err := ioutil.ReadAll(read)
	if err != nil {
		sugar.Log.Error(" Read all remote cid content is failed.Err:", err)
	}
	sugar.Log.Info(" remote file info :", string(remote1))
	//  update file.
	//  read local file info.
	//

	var defaltPath = path + "local"
	sugar.Log.Info(" Local file path :", defaltPath)
	sugar.Log.Info(" Exist local file ")

	lfile := Exist(defaltPath)
	sugar.Log.Info(" All cid info = local cid + _ + cid(update)")

	var upInfo string = string(remote1) + "_" + updateCid
	if !lfile {
		//create file
		sugar.Log.Info(" No find local file , and create it. ")

		_, err1 := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //open file
		if err1 != nil {
			sugar.Log.Errorf(" Create %s file is failed.Err:", err1)
		}

	}
	sugar.Log.Info(" Local file is exist,and open it. ")

	f1, err := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //open file
	if err != nil {
		sugar.Log.Errorf(" Open %s file is failed.Err:", err)
	}
	var all = string(remote1) + "_" + updateCid
	//upInfo = all
	sugar.Log.Info(" Upload All cid info,all : ", all)
	sugar.Log.Info(" Upload All cid info,upInfoCid: ", upInfo)
	_, err = f1.WriteString(all)
	if err != nil {
		sugar.Log.Error(" Write file is failed. Err:", err)
	}

	//	upload remote file to ipfs .
	sugar.Log.Info(" start upload allcid to ipfs . ")
	hash1, err := sh.Add(bytes.NewBufferString(upInfo))
	if err != nil {
		sugar.Log.Error(" upload file to ipfs is failed.Err: ", err)
	}
	sugar.Log.Info(" all ipfs hash :  ", hash1)
	// upload local update file to ipfs ,return a hash cid.
	// then need use ipns name publish -key=dbkey  to publish .
	ctx := context.Background()
	ksys, _ := sh.KeyList(ctx)
	sugar.Log.Info("  About all ipns key : ", ksys)
	// // fmt.Println(" keys 的 2集合 ：", ksys[2].Id)
	// // fmt.Println(" keys 的 2集合 ：", ksys[2].Name)
	sugar.Log.Info(" Look for the  db-key is exist in local keys array. ")
	var dbexist bool
	if len(ksys) > 0 {
		for _, v := range ksys {
			if v.Name == "dbkey" && v.Id == "k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v" {
				dbexist = true
				break
			}
		}
		if !dbexist {
			sugar.Log.Info(" Because the dbkey is inexist,then add it to local serct keys .")
			sugar.Log.Info(" dbkey path : ", path)
			postFormDataWithSingleFile(path)
		}
	}
	sugar.Log.Info(" Use ipns publish cid to public gateway.io ")
	//time duration
	t := time.Duration(time.Hour * 24)
	sugar.Log.Infof(" -- Excute ipns name publish -key=%s /ipfs/%s . --\n", "k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v", hash1)
	pubresp, err := sh.PublishWithDetails("/ipfs/"+hash1, "k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v", t, t, true)
	if err != nil {
		sugar.Log.Error(" PublishWithDetails is failed.Err: ", err)
	}
	sugar.Log.Info(" Pubresp := ", pubresp)
	sugar.Log.Info(" Off Line  Sync is Successful !!!! ")

}

// request ipns

func postFormDataWithSingleFile(path string) {
	sugar.Log.Info("  Start import dbkey.  ")
	sugar.Log.Info("  Import dbkey path :", path)
	client := http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	file, err := os.Open(path + "db-key")
	if err != nil {
		sugar.Log.Error("  Open dbkey path is failed.Err: ", err)
	}
	// file as key
	fileWrite, err := bodyWrite.CreateFormFile("file", "db-key")
	if err != nil {
		sugar.Log.Error("  CreateFormFile is failed.Err: ", err)
	}
	_, err = io.Copy(fileWrite, file)
	if err != nil {
		sugar.Log.Error(" Copy is failed.Err: ", err)
	}
	bodyWrite.Close() //will closed, 会将w.w.boundary刷写到w.writer中
	// create requet.
	contentType := bodyWrite.FormDataContentType()
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:5001/api/v0/key/import?arg=dbkey&ipns-base=base36", bodyBuf)
	if err != nil {
		sugar.Log.Error(" NewRequestpy is failed.Err: ", err)
	}
	// set request header.
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		sugar.Log.Error(" NewRequestpy is failed.Err: ", err)

	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sugar.Log.Error(" ReadAll resp result is failed.Err: ", err)
	}
	sugar.Log.Info(" Response restult: ", string(b))
	sugar.Log.Info(" Import dbkey is successful !!! ")
	defer file.Close()
}
