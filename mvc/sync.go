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
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"
	shell "github.com/ipfs/go-ipfs-api"
	ipfsCore "github.com/ipfs/go-ipfs/core"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/mr-tron/base58/base58"
)

func SyncUser(db *Sql, value string) error {
	sugar.Log.Info("---Start Sync User ---- ")
	var user SysUser
	err := json.Unmarshal([]byte(value), &user)
	if err != nil {
		sugar.Log.Error("Sync User Unmarshal is failed.Err:", err)
		return err
	}
	sugar.Log.Info("params:= ", user)
	l, e := FindIsExistUser(db, user)
	if e != nil {
		sugar.Log.Error("FindIsExistUser info is Failed.")
	}
	// l > 0 user is exist.
	if l > 0 {
		sugar.Log.Error("user is exist.")
		return errors.New(" User is already exist. ")
	}
	//inExist insert into sys_user.
	//create now time
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO sys_user values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.Err:", err)
		return err
	}
	//sid := strconv.FormatInt(user.Id, 10)
	res, err := stmt.Exec(user.Id, user.PeerId, user.Name, user.Phone, user.Sex, t, t, user.NickName, user.Img)
	if err != nil {
		sugar.Log.Error("Insert data to sys_user is failed.Err:", err)
		return err
	}
	c, _ := res.RowsAffected()
	if c == 0 {
		return errors.New(" Insert into sys_user is failed. ")
	}
	sugar.Log.Info("---Start Sync User End---- ")

	sugar.Log.Info("~~~~~   Insert into sys_user data is Successful ~~~~~~", c)
	return nil
}

// 文章

func SyncArticle(db *Sql, value string) error {
	sugar.Log.Info("---Start Sync  Article ---- ")

	var art vo.ArticleAddParams
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error("Marshal is failed.Err is ", err)
		return errors.New(" SyncArticle marshal is failed. ")
	}
	sugar.Log.Info("Marshal data is  ", art)
	// id := utils.SnowId()
	t := time.Now().Unix()
	stmt, err := db.DB.Prepare("INSERT INTO article values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		sugar.Log.Error("Insert into article table is failed.", err)
		return err
	}
	// sid := strconv.FormatInt(id, 10)
	stmt.QueryRow()
	res, err := stmt.Exec(art.Id, art.UserId, art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t, 0, 0, art.Title, art.Thumbnail, art.FileName, art.FileSize)
	if err != nil {
		sugar.Log.Error("Insert into article  is Failed.", err)
		return err
	}
	l, _ := res.RowsAffected()
	if l == 0 {
		return errors.New(" SyncArticle insert into article is failed. ")
	}
	sugar.Log.Info("---Start Sync  Article  End---- ")
	return nil
}

// 同步 文章 播放量

func SyncAticlePlay(db *Sql, value string) error {
	sugar.Log.Info("---Start Sync  AticlePlay ---- ")
	var dl Article
	var art vo.ArticlePlayAddParams
	//marshal request params.
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
	rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("Query data is failed.Err is ", err)
		return err
	}
	//scan data.
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Query scan data is failed.The err is ", err)
			return err
		}
		sugar.Log.Info("Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" Update is failed . ")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set play_num=? where id=?")
	if err != nil {
		sugar.Log.Error("Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.PlayNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Exec update is failed.The err is ", err)
		return err
	}
	//if affect equal zreo,it meant update is failed.
	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("RowsAffected  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Update  is failed.The err is ", err)
		return err
	}
	sugar.Log.Info("---Start Sync  AticlePlay  End ---- ")
	return nil
}

// 同步用户分享数量

func SyncArticleShareAdd(db *Sql, value string) error {
	sugar.Log.Info("---Start Sync ArticleShareAdd   ---- ")
	var dl Article
	var art vo.ArticlePlayAddParams
	//unmarshal request params.
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error(" Sync articleShare Add Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info("SyncArticleShareAdd Marshal data:", art)
	//Query whether the data exists.
	rows, err := db.DB.Query("select id,IFNULL(user_id,'null'),IFNULL(accesstory,'null'),IFNULL(accesstory_type,0),IFNULL(text,'null'),IFNULL(tag,'null'),IFNULL(ptime,0),IFNULL(play_num,0),IFNULL(share_num,0),IFNULL(title,'null'),IFNULL(thumbnail,'null'),IFNULL(file_name,'null'),IFNULL(file_size,0) from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error("SyncArticleShareAdd Query data is failed.Err is ", err)
		return err
	}
	//scan data => article.
	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Sync Query scan data is failed.Err: ", err)
			return err
		}
		sugar.Log.Info(" Query a entire data is : ", dl)
	}
	if dl.Id == "" {
		return errors.New(" Update is failed . ")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error(" Sync Update  data is failed.Err: ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Sync Update  is failed.Err: ", err)
		return err
	}
	//rowsAffect.
	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error("Sync Update  is failed.Err: ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error("Sync Update  is failed.Err: ", err)
		return err
	}
	sugar.Log.Info("---Start Sync ArticleShareAdd   End ---- ")
	return nil
}

// 同步用户注册
// 最开始写的  不会用。错误。
func SyncUserRegister(db *Sql, value string) error {
	//
	sugar.Log.Info("---Start   Sync   UserRegister    ---- ")

	var dl Article
	var art vo.ArticlePlayAddParams
	//unmarshal request params.
	err := json.Unmarshal([]byte(value), &art)
	if err != nil {
		sugar.Log.Error(" Sync Marshal is failed.Err is ", err)
		return err
	}
	sugar.Log.Info(" Marshal data is:", art)
	if err != nil {
		sugar.Log.Error(" Insert into article table is failed.", err)
		return err
	}
	//Query whether the data exists
	rows, err := db.DB.Query("select * from article where id=?", art.Id)
	if err != nil {
		sugar.Log.Error(" Query data is failed.Err is ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&dl.Id, &dl.UserId, &dl.Accesstory, &dl.AccesstoryType, &dl.Text, &dl.Tag, &dl.Ptime, &dl.PlayNum, &dl.ShareNum, &dl.Title, &dl.Thumbnail, &dl.FileName, &dl.FileSize)
		if err != nil {
			sugar.Log.Error("Sync Query scan data is failed.The err is ", err)
			return err
		}

		sugar.Log.Info("Sync Query a entire data is ", dl)
	}
	if dl.Id == "" {
		return errors.New(" Sync update is failed .")
	}
	//update play num + 1
	stmt, err := db.DB.Prepare("update article set share_num=? where id=?")
	if err != nil {
		sugar.Log.Error(" Sync Update  data is failed.The err is ", err)
		return err
	}
	res, err := stmt.Exec(int64(dl.ShareNum+1), art.Id)
	if err != nil {
		sugar.Log.Error("Sync Update  is failed.The err is ", err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		sugar.Log.Error(" Sync Update  is failed.The err is ", err)
		return err
	}
	if affect == 0 {
		sugar.Log.Error(" Sync Update  is failed.The err is ", err)
		return err
	}
	sugar.Log.Info("---Start   Sync   UserRegister   End ---- ")
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
	sugar.Log.Info("----  Start Sync Data ----")
	//Listening to the topic.
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("This is recover info:", err)
		}
	}()
	//The first step.
	topic := "/db-online-sync"
	sugar.Log.Info("Topic Name: ", topic)
	sugar.Log.Info("Subscrib Topic: ", topic)
	ctx := context.Background()
	sugar.Log.Info("Join Topic Room: ", topic)
	// join topic.
	// if topic is exist, use it (tp),otherwise join it.
	// ok true exist  false inexist.
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
	// start subscribe topic.
	sub, err := tp.Subscribe()
	if err != nil {
		sugar.Log.Error("subscribe failed.", err)
		return err
	}
	for {
		sugar.Log.Info("-----  Start Subscribe ------")
		data, err := sub.Next(ctx)
		if err != nil {
			sugar.Log.Error("subscribe failed.", err)
			continue
		}
		sugar.Log.Info("~~~  Recieve Data  ~~~~")
		msg := data.Message
		fromId := msg.From
		peerId := ipfsNode.Identity.String()
		sugar.Log.Info(" Recieve Data :", msg.Data)
		sugar.Log.Infof(" Recieve Data Type : %T\n", msg.Data)
		sugar.Log.Info(" Recieve FromId : ", string(fromId))
		sugar.Log.Info(" Local PeerId : ", peerId)
		//marshal recieve data.
		var recieve vo.SyncMsgParams
		err = json.Unmarshal(msg.Data, &recieve)
		if err != nil {
			sugar.Log.Error("Marshal recieve data is failed.Err:", err)
			continue
		}
		wayId := "12D3KooWDoBhdQwGT6oq2EG8rsduRCmyTZtHaBCowFZ7enwP4i8J"
		wayId2 := "12D3KooW9tijGFxQnN88QotdNfb77hGh3sXEiLc3NWPU4zWbN9WQ"
		wayId3 := "12D3KooWRxvZGzeMcAbxXuomztAwn344EkmiRusF7x5H3U4RtkNN"
		sugar.Log.Info("----Public gateway peerId1 ----:", wayId)
		sugar.Log.Info("----Public gateway peerId2 ----:", wayId2)
		sugar.Log.Info("----Public gateway peerId3 ----:", wayId3)
		sugar.Log.Info("----Encode base58 fromId ----")
		FromID := base58.Encode(fromId)
		sugar.Log.Info("----- Encode base58 fromId -----:", string(FromID))
		sugar.Log.Info("----- Judge FromId == or != wayId -----")
		if FromID == wayId || FromID == wayId2 || FromID == wayId3 {
			//if fromid == wayid ,then judge peerid ==reciece.fromid .
			//Satisfy one condition
			sugar.Log.Info(" FromId == wayId ")

			if peerId == recieve.FromId {
				sugar.Log.Info(" PeerId   !=   recieve.FromId")
				sugar.Log.Info(" PeerId :=", peerId)
				sugar.Log.Info(" recieve.FromId :=", recieve.FromId)
				sugar.Log.Info(" ~~~~  continue ~~~~ ")
				continue
			} else {
				sugar.Log.Info(" PeerId   ==   recieve.FromId")
				sugar.Log.Info(" PeerId :=", peerId)
				sugar.Log.Info(" recieve.FromId :=", recieve.FromId)
				sugar.Log.Info(" recieve.Method :=", recieve.Method)
				if recieve.Method == "receiveArticleAdd" {
					sugar.Log.Info("~~~  Start add  article   ~~~")
					//  add article into table.
					sugar.Log.Info("~~~ Because Method == receiveArticleAdd ~~~~")
					sugar.Log.Info(" recieve.Method :=", recieve.Method)
					//unmarshal params.
					var syn vo.SyncRecieveArticleParams
					err = json.Unmarshal(msg.Data, &syn)
					if err != nil {
						sugar.Log.Error("Sync marshal params is failed.Err:", err)
						continue
					}
					// string
					// marshal syn.data => userInfo.
					userInfo, err := json.Marshal(syn.Data)
					if err != nil {
						sugar.Log.Error("Sync marshal params is failed.Err:", err)
						continue
					}
					//start sync article.
					err = db.SyncArticle(string(userInfo))
					if err != nil {
						sugar.Log.Error("Sync marshal params is failed.Err:", err)
						continue
					}
					sugar.Log.Info("~~~  Sync add article is successful!  ~~~")
					sugar.Log.Info("~~~   Add  article  End ~~~")
				} else if recieve.Method == "receiveArticlePlayAdd" {
					sugar.Log.Info("~~~  Start ReceiveArticlePlayAdd   ~~~")
					//  add article into table.
					sugar.Log.Info("~~~ Because Method == receiveArticlePlayAdd ~~~~")
					sugar.Log.Info(" recieve.Method :=", recieve.Method)
					//unmarshal params.
					var syn vo.SyncRecievePlayParams
					err = json.Unmarshal(msg.Data, &syn)
					if err != nil {
						sugar.Log.Error("Sync marshal params is failed.Err:", err)
						continue
					}
					// string
					// marshal syn.data => userInfo.
					userInfo, err := json.Marshal(syn.Data)
					if err != nil {
						sugar.Log.Error("Sync marshal params is failed.Err:", err)
						continue
					}
					//sync articlplay
					err = db.SyncArticlePlay(string(userInfo))
					if err != nil {
						sugar.Log.Error("Sync article  add  play is failed.Err:", err)
						continue
					}
					sugar.Log.Info("~~~  Sync Article  Add Play is successful!  ~~~")
					sugar.Log.Info("~~~    Article  add  play  End   ~~~")
				} else if recieve.Method == "receiveArticleShareAdd" {
					sugar.Log.Info("~~~  Start receiveArticleShareAdd   ~~~")
					//  add article into table.
					sugar.Log.Info("~~~ Because Method == receiveArticleShareAdd ~~~~")
					sugar.Log.Info(" recieve.Method :=", recieve.Method)
					//unmarshal params.
					var syn vo.SyncRecievePlayParams
					err = json.Unmarshal(msg.Data, &syn)
					if err != nil {
						sugar.Log.Error("Sync marshal params is failed.Err:", err)
						continue
					}
					// string
					// marshal syn.data => userInfo.
					userInfo, err := json.Marshal(syn.Data)
					if err != nil {
						sugar.Log.Error("Marshal params is failed.Err:", err)
						continue
					}
					// start sync ArticleShareAdd
					err = db.SyncArticleShareAdd(string(userInfo))
					if err != nil {
						sugar.Log.Error("-Sync articleshareadd is failed.Err:", err)
						continue
					}
					sugar.Log.Info("~~~  Sync articleshareadd is successful!  ~~~")
					sugar.Log.Info("~~~   Add  articleshareadd  End ~~~")
				} else if recieve.Method == "receiveUserRegister" {
					sugar.Log.Info("~~~  Start receiveUserRegister   ~~~")
					//  add article into table.
					sugar.Log.Info("~~~ Because Method == receiveUserRegister ~~~~")
					sugar.Log.Info(" recieve.Method :=", recieve.Method)
					//unmarshal params.
					var syn vo.SyncRecieveUsesrParams
					err = json.Unmarshal(msg.Data, &syn)
					if err != nil {
						sugar.Log.Error("Marshal params is failed.Err:", err)
						continue
					}
					// string
					// marshal syn.data => userInfo.
					userInfo, err := json.Marshal(syn.Data)
					if err != nil {
						sugar.Log.Error("Marshal params is failed.Err:", err)
						continue
					}
					//start sync user.
					err = db.SyncUser(string(userInfo))
					if err != nil {
						sugar.Log.Error("Sync user is failed.Err:", err)
						continue
					}
					sugar.Log.Info("~~~  Sync user is successful!  ~~~")
					sugar.Log.Info("~~~   Add  user  End ~~~")
				} else {
					sugar.Log.Info("~~~~~  No ~~~~~ ")
					sugar.Log.Info("~~~~~  Continue ~~~~~ ")
					sugar.Log.Info("~~~~~  Don't match ~~~~~ ")
					continue
				}
			}
		} else {
			sugar.Log.Info("~~~~~  No ~~~~~ ")
			sugar.Log.Info("~~~~~  Continue ~~~~~ ")
			sugar.Log.Info("~~~~~  Don't match ~~~~~ ")
			continue
		}

	}
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
			sugar.Log.Info(" ~~~~~~~~~ Capture the panic ~~~~~~~~~~~~Err: ", err)
		} else {
			sugar.Log.Info("~~~~~~~~~~~~~~~   Normal ~~~~~~~~~~~~")
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
	result, err := sh.Resolve("k2k4r8nzb4pahkhkvojnx3rtd0j9b432wjdcdjkogau637ix3zhmv1pe")
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
	sugar.Log.Info(" Read local file content : ", string(local))
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
		//removeDuplication
		remoteStr1 := strings.Split(string(remote1), "_")
		localStr1 := strings.Split(string(local), "_")

		dupremote := RemoveDuplicationArray(remoteStr1)
		duplocal := RemoveDuplicationArray(localStr1)
		sugar.Log.Info(" Duplication remote  array: ", dupremote)
		sugar.Log.Info(" Duplication local  array: ", duplocal)

		diff := difference(duplocal, dupremote)

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
			defer f1.Close()
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

		dupremoteStr := SplitArray(dupremote)
		sugar.Log.Info(" DupremoteStr data : ", dupremoteStr)
		_, err = local_f.WriteString(string(dupremoteStr))
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
	}
	//ipns
	sugar.Log.Info(" Start upload cid to gateway.io ipns. ")
	UploadFile(path, hash)
	sugar.Log.Info(" Because local  =====   remote.  ")
}

//split

func SplitArray(a []string) string {
	var result string
	for k, v := range a {
		if k == len(a)-1 {
			result += v
		} else {
			result += v + "_"
		}
	}
	return result
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

// find same data in array.
func RemoveDuplicationArray(arr []string) []string {
	set := make(map[string]struct{}, len(arr))
	j := 0
	for _, v := range arr {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		arr[j] = v
		j++
	}
	return arr[:j]
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
	// var upInfo string = string(remote1) + "_" + updateCid
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
	//
	remoteStr1 := strings.Split(string(remote1), "_")
	reduplication := RemoveDuplicationArray(remoteStr1)
	remoteresp := SplitArray(reduplication)
	sugar.Log.Info(" Duplicate removal remoteresp : ", remoteresp)
	sugar.Log.Info(" Duplicate removal len : ", len(reduplication))
	var all = remoteresp + "_" + updateCid
	// dup
	sugar.Log.Info(" This is  all : ", all)

	//upInfo = all
	sugar.Log.Info(" Upload All cid info,all : ", all)
	//sugar.Log.Info(" Upload All cid info,upInfoCid: ", upInfo)
	_, err = f1.WriteString(all)
	if err != nil {
		sugar.Log.Error(" Write file is failed. Err:", err)
	}
	//	upload remote file to ipfs .
	sugar.Log.Info(" start upload allcid to ipfs . ")
	hash1, err := sh.Add(bytes.NewBufferString(all))
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
			if v.Name == "dbkey" && v.Id == "k2k4r8nzb4pahkhkvojnx3rtd0j9b432wjdcdjkogau637ix3zhmv1pe" {
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
	sugar.Log.Infof(" -- Excute ipns name publish -key=%s /ipfs/%s . --\n", "k2k4r8nzb4pahkhkvojnx3rtd0j9b432wjdcdjkogau637ix3zhmv1pe", hash1)
	pubresp, err := sh.PublishWithDetails("/ipfs/"+hash1, "k2k4r8nzb4pahkhkvojnx3rtd0j9b432wjdcdjkogau637ix3zhmv1pe", t, t, true)
	if err != nil {
		sugar.Log.Error(" PublishWithDetails is failed.Err: ", err)
	}
	sugar.Log.Info(" Pubresp := ", pubresp)
	sugar.Log.Info("~~~~~ Off Line  Sync is Successful !!!! ~~~~~")
}

// request ipns

func postFormDataWithSingleFile(path string) {
	sugar.Log.Info("  Start import dbkey.  ")
	sugar.Log.Info("  Import dbkey path :", path)
	client := http.Client{}
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)
	file, err := os.Open(path + "db-key.key")
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
	bodyWrite.Close() //will closed, 会将w.w.boundary copy => w.writer
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
