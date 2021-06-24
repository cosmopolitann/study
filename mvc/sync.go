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

//离线同步数据。
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
func OffLineSyncData(db *Sql, path string) {
	//
	// sh = shell.NewShell("localhost:5001")
	// hash := "QmaZMLejnjNKex6Nrs2RGLC8n7NvWQP8RFPn2dLs2XviYb"
	// err := sh.Get(hash, "./output")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //cat
	// read, err := sh.Cat(hash)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// body, err := ioutil.ReadAll(read)
	// log.Println(string(body))

	// 创建 local 文件。
	sugar.Log.Info("--- Start excute offline task ---")
	var defaltPath = path + "local"
	sugar.Log.Info(" Local Path :", defaltPath)
	b := Exist(defaltPath)
	if !b {
		//创建
		sugar.Log.Info(" Local File is exist and create local file. ")
		_, err1 := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
		if err1 != nil {
			sugar.Log.Error(" Create Local File Is Failed.Err: ", err1)
		}
	}

	//拉取 cid 文件

	//按行读文件
	// var remote123 = "/Users/apple/winter/offline/remote"
	// ipfs 拉取 文件cid  固定位置
	sh = shell.NewShell("localhost:5001")
	// hash := "QmYntasS515q9oF2LC6Boka2aWAGs1EHnSdRfQzBYipH8j"
	//  解析 远程 remote ipns 的 cid 数据  并且 拉取内容
	sugar.Log.Info(" ---  Start resolve remote ipns data.  ---")

	// result, err := sh.Resolve("k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v")
	// RemoteIpnsAddr
	sugar.Log.Info(" Ipns Addr: ", RemoteIpnsAddr)

	result, err := sh.Resolve(RemoteIpnsAddr)

	if err != nil {
		sugar.Log.Error(" Ipns Addr resolve is failed. Err:", err)
	}
	sugar.Log.Info("The result what ipfs cat remote cid. ", result)

	hash := result
	sugar.Log.Info("Remote ipns addr cid : ", hash)

	// err := sh.Get(hash, remote123)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// 将远程文件 cid 拉取到本地之后  比对 差异 然后拉取 对应 cid  读取 文件 遍历 sql

	//读出  本地 文件
	log.Println(" ----- 开始  读取 本地 文件 local  信息 ----- ")
	sugar.Log.Info(" Read local file content. ")
	local, err := ioutil.ReadFile(path + "local") // just pass the file name
	if err != nil {
		sugar.Log.Error(" Read local file content is failed. ", err)

	}
	sugar.Log.Info(" Read local file content. ", string(local))

	//读出  远程 文件
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
				sugar.Log.Infof(" Data type for each line is %T .\n", line)
				// exec sql read cidPath file by line.
				sugar.Log.Info(" Start excute sql  by read cidpath file content. ", line)
				stmt, err := db.DB.Prepare(string(line))
				if err != nil {
					sugar.Log.Error("Insert data into table is failed.", err)
					continue
				}
				res, err := stmt.Exec()
				if err != nil {
					sugar.Log.Error("Insert data into  is Failed.", err)
					continue
				}
				l, err := res.RowsAffected()
				if l == 0 {
					sugar.Log.Error("Excute sql is failed.Err:", err)
					continue
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
					sugar.Log.Error(" Delete cidPath file is successful !!! ", cidPath)
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

		fmt.Println(" 开始 删除 remote 文件")
		existed := true
		if _, err := os.Stat(defaltPath + "remote"); os.IsNotExist(err) {
			existed = false
		}
		if existed {
			err := os.Remove(defaltPath + "remote")

			if err != nil {
				fmt.Println(" 删除失败 cidPath 文件 ", defaltPath+"remote")
				// 删除失败
			} else {
				fmt.Println(" 删除成功 cidPath ", defaltPath+"remote")
				// 删除成功
			}
		}
		//  删除文件

		sugar.Log.Info("-----------------------   执行sql 成功  ---------------")

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

		// 	}

		// }

	}

	//
	fmt.Println(" 开始 执行  更新 本地 数据  到 ipns ")
	UploadFile(path, hash)

	fmt.Println("远程和本地相等，不执行任何操作，直接返回。")

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

//  本地更新文件

func UploadFile(path string, hash string) {
	//  解析 k5 id  然后 拉取对应的 remote 数据

	var updateCid string
	b := Exist(path + "update")
	if !b {
		//创建文件
		_, err1 := os.OpenFile(path+"update", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666) //打开文件
		if err1 != nil {
			fmt.Println("创建失败")
		}

		// } else {
	} else {
		bytes1, err := ioutil.ReadFile(path + "update")

		if err != nil {
			fmt.Println("读取内容失败", err)

		}
		// 上传 ipfs
		hash_local, err := sh.Add(bytes.NewBufferString(string(bytes1)))
		if err != nil {
			fmt.Println("上传ipfs时错误：", err)
		}
		fmt.Println("这是上传的时候 hash_local == ", hash_local)
		updateCid = hash_local
	}
	// 默认的文件 hash
	// hash := "QmYntasS515q9oF2LC6Boka2aWAGs1EHnSdRfQzBYipH8j"
	// hash := result
	fmt.Println("  测试 更新  文件 上传")
	//读出  远程 文件
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	remote1, err := ioutil.ReadAll(read)
	fmt.Println("  这是 读出的 remote 远程文件的信息内容：", remote1)
	//  检查本地是否有 更新文件
	//  读出本地 local 文件 信息内容
	var defaltPath = path + "local"

	lfile := Exist(defaltPath)

	var upInfo string = string(remote1) + "_" + updateCid
	if lfile == false {
		//创建
		_, err1 := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
		if err1 != nil {
			fmt.Println("创建失败")
		}
		fmt.Println("  ----- 本地 local 文件 存在  ----")

	} else {
		f1, err := os.OpenFile(defaltPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件

		var all = string(remote1) + "_" + updateCid
		//upInfo = all
		fmt.Println("  ----- 上传的 信息 upInfo  ----", upInfo)

		_, err = f1.WriteString(all)
		if err != nil {
			fmt.Println(" 写入 local 文件 错误：", err)
		}
	}

	//	在上传remote 文件到 ipfs.
	fmt.Println("  ----- 开始上传 remote 文件 到 ipfs   ----", upInfo)
	hash1, err := sh.Add(bytes.NewBufferString(upInfo))
	if err != nil {
		fmt.Println("上传ipfs时错误：", err)
	}
	fmt.Println("这是上传的时候 hash1 == ", hash1)
	// 上传本地 update 文件到 ipfs 生成cid
	// 需要将 remote cid  ipns 到 一个地方。
	ctx := context.Background()
	ksys, _ := sh.KeyList(ctx)
	fmt.Println(" keys 的 集合 ：", ksys)
	// // fmt.Println(" keys 的 2集合 ：", ksys[2].Id)
	// // fmt.Println(" keys 的 2集合 ：", ksys[2].Name)

	// fmt.Println(" keys 的 1集合 ：", ksys[0].Id)
	// fmt.Println(" keys 的 1集合 ：", ksys[0].Name)
	// // puberr := sh.Publish("", "/ipfs/QmYSctvKQMjZ51RybBcXzht2GRME6aXvvgeBUV8QFJLoBr")
	// // if puberr != nil {
	// // 	fmt.Println(" pubsub ipns 失败 =", puberr)
	// // }

	fmt.Println("----完成 ---- =")
	//查看 本地 是否有 dbkey 这个秘钥 如果没有 就 加入 如果有 就直接上传
	fmt.Println(" keys 的 集合 ：", ksys)
	var dbexist bool
	if len(ksys) > 0 {
		for _, v := range ksys {
			if v.Name == "dbkey" && v.Id == "k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v" {
				dbexist = true
				break
			}
		}
		if dbexist == false {
			fmt.Println(" ----- 因为  里面  没有 dbkey  所以 添加 秘钥 -----")
			postFormDataWithSingleFile(path)
		}
	}

	// result, err := sh.Resolve("k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v")
	// if err != nil {
	// 	fmt.Println(" 解析 k5 id 失败 =", err)
	// }
	// fmt.Println(" 解析 k5 id 结果 =", result)

	t := time.Duration(time.Hour * 24)
	fmt.Println("-----  开始 执行 pubsbu ipns -----")
	pubresp, err := sh.PublishWithDetails("/ipfs/"+hash1, "k51qzi5uqu5dl2hdjuvu5mqlxuvezwe5wbedi6uh7dgu1uiv61vh4p4b71b17v", t, t, true)
	if err != nil {
		fmt.Println(" pubsub content 失败  =", err)
	}

	fmt.Println("pubresp =", pubresp)

	//http 请求 ipns

}

// 请求 ipns

func postFormDataWithSingleFile(path string) {
	fmt.Println("------  开始 导入 dbkey ------")
	client := http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	//路径  传进来。。
	//todo

	file, err := os.Open(path + "db-key")
	defer file.Close()
	if err != nil {
		log.Println("err")
	}
	// file 为key
	fileWrite, err := bodyWrite.CreateFormFile("file", "db-key")
	_, err = io.Copy(fileWrite, file)
	if err != nil {
		log.Println("err")
	}
	bodyWrite.Close() //要关闭，会将w.w.boundary刷写到w.writer中
	// 创建请求
	contentType := bodyWrite.FormDataContentType()
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:5001/api/v0/key/import?arg=dbkey&ipns-base=base36", bodyBuf)
	if err != nil {
		log.Println("err")
	}
	// 设置头
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err")
	}
	fmt.Println(string(b))
	fmt.Println("------  开始 导入 dbkey  成功------")

}
