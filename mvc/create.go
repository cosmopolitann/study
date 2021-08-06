package mvc

import (
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/sugar"
	"github.com/cosmopolitann/clouddb/vo"

	ipfsCore "github.com/ipfs/go-ipfs/core"
)

//|-------------------------------------------------|
// 初始化数据库

func (db *Sql) Ping() error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> Ping, Err: %v", err)
		}
	}()

	err := db.DB.Ping()
	if err != nil {
		sugar.Log.Error("Ping is Failed.", err)
	} //|
	return err //|
}

//---------------------------------------------------|

/*
------------------------------------------------------
|                        User                        |
------------------------------------------------------
*/

//  用户注册

func (db *Sql) UserRegister(ipfsNode *ipfsCore.IpfsNode, user string, path string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> UserRegister, Err: %v", err)
		}
	}()

	data, err := AddUser(ipfsNode, db, user, path)
	//返回封装成方法
	// 返回的时候 要改东西
	if err != nil {
		return vo.ResponseErrorMsg(400, err.Error())
	}
	return vo.ResponseSuccess(data)
}

func (db *Sql) AddUserTest(user string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> AddUserTest, Err: %v", err)
		}
	}()

	data, err := AddUserTest(db, user)
	//返回封装成方法
	// 返回的时候 要改东西
	if err != nil {
		return vo.ResponseErrorMsg(400, err.Error())
	}
	return vo.ResponseSuccess(data)
}

//  用户注销

func (db *Sql) UserLoginOut(user string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> UserLoginOut, Err: %v", err)
		}
	}()

	e := UserDel(db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

//   用户登录

func (db *Sql) UserLogin(user string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> UserLogin, Err: %v", err)
		}
	}()

	token, e := UserLogin(db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(token)
}

//  用户查询

func (db *Sql) UserQuery(user string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> UserQuery, Err: %v", err)
		}
	}()

	data, e := UserQuery(db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(data)
}

// 其他 用户 信息 查询

func (db *Sql) OtherUserQuery(user string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> OtherUserQuery, Err: %v", err)
		}
	}()

	data, e := OtherUserQuery(db, user)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(data)
}

//用户 更新

func (db *Sql) UserUpdate(ipfsNode *ipfsCore.IpfsNode, user string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> UserUpdate, Err: %v", err)
		}
	}()

	e := UserUpdate(ipfsNode, db, user)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

/*
------------------------------------------------------
|                     User  End                      |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                     Friend                         |
------------------------------------------------------
*/

// 检测好友是否在线

func (db *Sql) FriendCheckOnline(ipfsNode *ipfsCore.IpfsNode, data string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> FriendCheckOnline, Err: %v", err)
		}
	}()

	e := FriendCheckOnline(ipfsNode, data)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

// 设置好友昵称
func (db *Sql) FriendUpdateNickname(data string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> FriendUpdateNickname, Err: %v", err)
		}
	}()

	e := FriendUpdateNickname(db, data)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

/*
------------------------------------------------------
|                     Friend  End                      |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                       Cloud                        |
------------------------------------------------------
*/

//  添加文件

func (db *Sql) AddFile(fInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> AddFile, Err: %v", err)
		}
	}()

	fileId, e := AddFile(db, fInfo)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(fileId)
}

//   添加文件夹

func (db *Sql) AddFolder(fInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> AddFolder, Err: %v", err)
		}
	}()

	e := AddFolder(db, fInfo)

	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

// 删除文件

func (db *Sql) DeleteOneFile(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> DeleteOneFile, Err: %v", err)
		}
	}()

	e := DeleteOneFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

//重命名

func (db *Sql) FileRename(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> FileRename, Err: %v", err)
		}
	}()

	e := CloudFileRename(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

// 获取文件层级列表

func (db *Sql) FileList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> FileList, Err: %v", err)
		}
	}()

	data, e := CloudFileList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 获取文件夹层级列表

func (db *Sql) FolderList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> FolderList, Err: %v", err)
		}
	}()

	data, e := CloudFolderList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 上传记录传输

func (db *Sql) TransferAdd(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> TransferAdd, Err: %v", err)
		}
	}()

	e := DownLoadFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

//  根据文件进行分类

func (db *Sql) FileCategory(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> FileCategory, Err: %v", err)
		}
	}()

	data, e := CloudFileCategory(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

//  删除传输记录

func (db *Sql) TransferDel(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> TransferDel, Err: %v", err)
		}
	}()

	e := TransferDel(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

//  传输列表

func (db *Sql) TransferList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> TransferList, Err: %v", err)
		}
	}()

	data, e := TransferList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

//  下载列表

func (db *Sql) DownloadList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> DownloadList, Err: %v", err)
		}
	}()

	data, e := DownloadList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

//  复制文件

func (db *Sql) CopyFile(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> CopyFile, Err: %v", err)
		}
	}()

	e := CopyFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

//  移动文件

func (db *Sql) MoveFile(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> MoveFile, Err: %v", err)
		}
	}()

	e := MoveFile(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

//  删除文件

func (db *Sql) DeleteAll(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> DeleteAll, Err: %v", err)
		}
	}()

	e := Delete(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

//delete

func (db *Sql) CloudFindList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> CloudFindList, Err: %v", err)
		}
	}()

	result, e := CloudFindList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(result)
}

//查询

func (db *Sql) CloudSearch(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> CloudSearch, Err: %v", err)
		}
	}()

	result, e := Search(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(result)
}

/*
------------------------------------------------------
|                    Cloud  End                      |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                     Article                        |
------------------------------------------------------
*/

//  添加 朋友圈文章

func (db *Sql) ArticleAdd(ipfsNode *ipfsCore.IpfsNode, dInfo, path string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleAdd, Err: %v", err)
		}
	}()

	e := AddArticle(ipfsNode, db, dInfo, path)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

//  查找文章详情

func (db *Sql) ArticleList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleList, Err: %v", err)
		}
	}()

	data, e := ArticleList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

//  用户文章列表
func (db *Sql) ArticleListUser(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleListUser, Err: %v", err)
		}
	}()

	data, e := ArticleListUser(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 文章列表分类

func (db *Sql) ArticleCategory(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleCategory, Err: %v", err)
		}
	}()

	data, e := ArticleCategory(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	//
	//resmodel := vo.BuildResp()
	//resmodel.Count=11
	//resmodel.Data=data
	return vo.ResponseSuccess(data)
}

// 文章增加播放次数

func (db *Sql) ArticlePlayAdd(ipfsNode *ipfsCore.IpfsNode, dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticlePlayAdd, Err: %v", err)
		}
	}()

	e := ArticlePlayAdd(ipfsNode, db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

// 增加播放量

func (db *Sql) ArticleShareAdd(ipfsNode *ipfsCore.IpfsNode, dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleShareAdd, Err: %v", err)
		}
	}()

	e := ArticleShareAdd(ipfsNode, db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

//  朋友圈 点赞

func (db *Sql) ArticleGiveLike(ipfsNode *ipfsCore.IpfsNode, dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleGiveLike, Err: %v", err)
		}
	}()

	e := AddArticleLike(ipfsNode, db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

// 取消点赞

func (db *Sql) ArticleCancelLike(ipfsNode *ipfsCore.IpfsNode, dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleCancelLike, Err: %v", err)
		}
	}()

	data, e := ArticleCancelLike(ipfsNode, db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 获取文章详情

func (db *Sql) ArticleQuery(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleQuery, Err: %v", err)
		}
	}()

	data, e := ArticleQuery(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 文章查询

func (db *Sql) ArticleSearch(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleSearch, Err: %v", err)
		}
	}()

	data, e := ARticleSearch(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 我发布的 文章信息

func (db *Sql) ArticleAboutMe(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleAboutMe, Err: %v", err)
		}
	}()

	data, e := ArticleAboutMe(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

//

func (db *Sql) ArticleSearchCagetory(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleSearchCagetory, Err: %v", err)
		}
	}()

	data, e := ArticleSearchCagetory(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 推荐  待定  随机10条
func (db *Sql) ArticleRecommend(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleRecommend, Err: %v", err)
		}
	}()

	data, e := ArticleRecommend(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}
func (db *Sql) ArticleRecommendLimitTenData(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleRecommendLimitTenData, Err: %v", err)
		}
	}()

	data, e := ArticleRecommendLimitTenData(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}
func (db *Sql) ArticleAddTest(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ArticleAddTest, Err: %v", err)
		}
	}()

	e := ArticleAddTest(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess()
}

/*
------------------------------------------------------
|                 Article  End                       |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                       Chat                         |
------------------------------------------------------
*/

// //保存消息

// func (db *Sql) ChatAddMsg(dInfo string) string {

// 	e := AddChatMsg(db, dInfo)
// 	if e != nil {
// 		return vo.ResponseErrorMsg(400, e.Error())
// 	}

// 	return vo.ResponseSuccess()
// }

//获取 消息 分页

func (db *Sql) ChatMsgList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatMsgList, Err: %v", err)
		}
	}()

	data, e := ChatMsgList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}

	return vo.ResponseSuccess(data)
}

// 删除消息

func (db *Sql) ChatMsgDel(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatMsgDel, Err: %v", err)
		}
	}()

	e := ChatMsgDel(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

// // 消息记录 新增

// func (db *Sql) ChatRecordAdd(dInfo string) string {

// 	recordId, e := ChatRecordAdd(db, dInfo)
// 	if e != nil {
// 		return vo.ResponseErrorMsg(400, e.Error())
// 	}
// 	return vo.ResponseSuccess(recordId)
// }

// 获取消息记录列表

func (db *Sql) ChatRecordList(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatRecordList, Err: %v", err)
		}
	}()

	data, e := ChatRecordList(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(data)
}

//  删除记录

func (db *Sql) ChatRecordDel(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatRecordDel, Err: %v", err)
		}
	}()

	e := ChatRecordDel(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

// //  撤回消息

// func (db *Sql) ChatMsgWithDraw(dInfo string) string {

// 	e := ChatWithDraw(db, dInfo)
// 	if e != nil {
// 		return vo.ResponseErrorMsg(400, e.Error())
// 	}
// 	return vo.ResponseSuccess()
// }

// //  聊天记录

// func (db *Sql) AddChatRecord(dInfo string) string {

// 	e := AddChatRecord(db, dInfo)
// 	if e != nil {
// 		return vo.ResponseErrorMsg(400, e.Error())
// 	}

// 	return vo.ResponseSuccess()
// }

// ChatCreateRecord  创建会话
func (db *Sql) ChatCreateRecord(ipfsNode *ipfsCore.IpfsNode, msg string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatCreateRecord, Err: %v", err)
		}
	}()

	data, err := ChatCreateRecord(ipfsNode, db, msg)
	if err != nil {
		return vo.ResponseErrorMsg(400, err.Error())
	}
	return vo.ResponseSuccess(data)
}

// // ChatCreateRecord  重命名会话
// func (db *Sql) ChatRenameRecord(ipfsNode *ipfsCore.IpfsNode, msg string) string {

// 	err := ChatRenameRecord(ipfsNode, db, msg)
// 	if err != nil {
// 		return vo.ResponseErrorMsg(400, err.Error())
// 	}
// 	return vo.ResponseSuccess()
// }

// ChatSendMsg  发送消息
func (db *Sql) ChatSendMsg(ipfsNode *ipfsCore.IpfsNode, msg string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatSendMsg, Err: %v", err)
		}
	}()

	data, err := ChatSendMsg(ipfsNode, db, msg)
	if err != nil {
		return vo.ResponseErrorMsg(400, err.Error())
	}
	return vo.ResponseSuccess(data)
}

// ChatReadMsg  已读消息
func (db *Sql) ChatReadMsg(msg string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatReadMsg, Err: %v", err)
		}
	}()

	err := ChatReadMsg(db, msg)
	if err != nil {
		return vo.ResponseErrorMsg(400, err.Error())
	}
	return vo.ResponseSuccess()
}

// ChatWithdrawMsg  撤回消息
func (db *Sql) ChatWithdrawMsg(ipfsNode *ipfsCore.IpfsNode, msg string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatWithdrawMsg, Err: %v", err)
		}
	}()

	err := ChatWithdrawMsg(ipfsNode, db, msg)
	if err != nil {
		return vo.ResponseErrorMsg(400, err.Error())
	}
	return vo.ResponseSuccess()
}

// ChatListenMsgBlocked  监听消息 阻塞式
func (db *Sql) ChatListenMsgBlocked(ipfsNode *ipfsCore.IpfsNode, token string, clh vo.ChatListenHandler) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatListenMsgBlocked, Err: %v", err)
		}
	}()

	err := ChatListenMsgBlocked(ipfsNode, db, token, clh)

	return err
}

// ChatListenMsgUpdateUser  更新当前用户
func (db *Sql) ChatListenMsgUpdateUser(ipfsNode *ipfsCore.IpfsNode, token string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ChatListenMsgUpdateUser, Err: %v", err)
		}
	}()

	err := ChatListenMsgUpdateUser(ipfsNode, token)

	return err
}

/*
------------------------------------------------------
|                    Chat  End                       |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                       Sync                         |
------------------------------------------------------
*/

//同步 User表 数据

func (db *Sql) SyncUser(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncUser, Err: %v", err)
		}
	}()

	e := SyncUser(db, dInfo)
	return e
}

//

// 同步文章点赞    article_like表

// 同步 文章 取消点赞  article_like表

func (db *Sql) SyncArticleCancelLike(dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncArticleCancelLike, Err: %v", err)
		}
	}()

	e := SyncAticlePlay(db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

//----------------------------------------------------------

//  同步 用户信息

func (db *Sql) SyncUserRegister(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncUserRegister, Err: %v", err)
		}
	}()

	e := SyncUserRegister(db, dInfo)
	return e
}

// 同步 article  表数据

func (db *Sql) SyncArticle(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncArticle, Err: %v", err)
		}
	}()

	e := SyncArticle(db, dInfo)
	return e
}

// 同步 article play  表数据

func (db *Sql) SyncArticlePlay(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncArticlePlay, Err: %v", err)
		}
	}()

	e := SyncAticlePlay(db, dInfo)

	return e
}

// 同步 文章 分享

func (db *Sql) SyncArticleShareAdd(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncArticleShareAdd, Err: %v", err)
		}
	}()

	e := SyncArticleShareAdd(db, dInfo)
	return e
}

// 同步 点赞
func (db *Sql) SyncArticleLike(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncArticleLike, Err: %v", err)
		}
	}()

	e := SyncArticleLike(db, dInfo)
	return e
}
func (db *Sql) SyncArticleCancelLikee(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncArticleCancelLikee, Err: %v", err)
		}
	}()

	e := SyncArticleCancelLike(db, dInfo)
	return e
}

//同步用户更新
func (db *Sql) SyncUserUpdate(dInfo string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncUserUpdate, Err: %v", err)
		}
	}()

	e := SyncUserUpdate(db, dInfo)
	return e
}

//  同步数据

func (db *Sql) SyncData(ipfsNode *ipfsCore.IpfsNode, dInfo string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncData, Err: %v", err)
		}
	}()

	sugar.Log.Info("---- Start  OnLine   Sync  ------")

	e := SyncTopicData(ipfsNode, db, dInfo)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess()
}

// 同步 聊天 用户 数据

func (db *Sql) SyncQueryAllData(dInfo string, path string) string {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncQueryAllData, Err: %v", err)
		}
	}()

	sugar.Log.Info("----   Start  QueryAll data     ------")

	e, data := SyncQueryAllData(dInfo, db, path)
	if e != nil {
		return vo.ResponseErrorMsg(400, e.Error())
	}
	return vo.ResponseSuccess(data)
}

func (db *Sql) SyncDatabaseMigration(token, path, cid string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> SyncDatabaseMigration, Err: %v", err)
		}
	}()

	sugar.Log.Info("----   Start  DatabaseMigration data    ------")
	e := SyncDatabaseMigration(token, path, cid, db)
	return e
}

//---------------------------------------------

/*
------------------------------------------------------
|                    Sync   End                      |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                   OffLine     Sync                 |
------------------------------------------------------
*/

func (db *Sql) OfflineSync(ipfsNode *ipfsCore.IpfsNode, path string) error {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> OfflineSync, Err: %v", err)
		}
	}()

	sugar.Log.Info("---- Start OffLine Sync  ------")
	// i := 0
	// c := cron.New()
	// // spec := "*/5 * * * * ?"
	// spec := "0 */5 * * * ?"
	// c.AddFunc(spec, func() {
	// 	//0 */1 * * * ?
	// 	log.Println("cron running:")
	// 	log.Println("path :", path)

	// 	OffLineSyncData(db, path)

	// })
	// c.Start()
	// select {}
	err := OffLineSyncData(db, path, ipfsNode)
	// if err != nil {
	// 	return vo.ResponseErrorMsg(400, err.Error())
	// }
	// return vo.ResponseSuccess()
	return err
}

/*
------------------------------------------------------
|             OffLine     Sync   End                 |
------------------------------------------------------
*/

/*
------------------------------------------------------
|                       Other                        |
------------------------------------------------------
*/

//convert

func ConvertString(value string, t interface{}) (res map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> ConvertString, Err: %v", err)
		}
	}()

	json.Unmarshal([]byte(value), &t)
	fmt.Println(" 这是 获得的结果 ", t)
	fmt.Printf(" 这是 获得的结果  %T\n", t)
	return t.(map[string]interface{})
}

// db upgrade
func (db *Sql) DbUpgrade(dbv string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			sugar.Log.Errorf("Panic -> DbUpgrade, Err: %v", err)
		}
	}()

	sugar.Log.Info("---- Start DbUpgrade   ------")

	newversion, err := DbUpgrade(db, dbv)
	// if err != nil {
	// 	return vo.ResponseErrorMsg(400, err.Error())
	// }
	// return vo.ResponseSuccess()
	return newversion, err
}
