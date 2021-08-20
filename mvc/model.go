package mvc

//sys_user

type SysUser struct {
	Id       string `json:"id"`       //id
	PeerId   string `json:"peerId"`   //节点id
	Name     string `json:"name"`     //用户名字
	Phone    string `json:"phone"`    //手机号
	Sex      int64  `json:"sex"`      //性别 0 未知  1 男  2 女
	NickName string `json:"nickName"` //昵称
	Ptime    int64  `json:"-"`        //时间
	Utime    int64  `json:"-"`        //更新时间
	Img      string `json:"img"`      //头像
	Role     string `json:"role"`     //角色

}

//cloud_file
type File struct {
	Id        string `json:"id"`        //id
	UserId    string `json:"userId"`    //用户userid
	FileName  string `json:"fileName"`  //文件名字
	ParentId  string `json:"parentId"`  //父id
	FileCid   string `json:"fileCid"`   //文件cid
	FileSize  int64  `json:"fileSize"`  //文件大小
	FileType  int64  `json:"fileType"`  //文件类型
	IsFolder  int64  `json:"isFolder"`  //是否是文件or 文件夹  0文件 1文件夹
	Ptime     int64  `json:"ptime"`     //时间
	Thumbnail string `json:"thumbnail"` //缩略图
	Duration  int64  `json:"duration"`  //时长
	Width     string `json:"width"`     //宽
	Height    string `json:"height"`    //高
}

//DownLoadList
type DownLoad struct {
	Id             string `json:"id"`             //id
	UserId         string `json:"userId"`         //用户uersid
	FileName       string `json:"fileName"`       //文件名字
	Ptime          int64  `json:"ptime"`          //时间
	FileCid        string `json:"fileCid"`        //文件cid
	FileSize       int64  `json:"fileSize"`       //文件大小
	DownPath       string `json:"downPath"`       //下载路径
	FileType       int64  `json:"fileType"`       //文件类型
	TransferType   int64  `json:"transferType"`   //传输类型 1 上传 2 下载
	UploadParentId string `json:"uploadParentId"` //下载父id
	UploadFileId   string `json:"uploadFileId"`   //下载文件的id
}

//
//DownLoadList
type TransferDownLoadParams struct {
	Id             string `json:"id"`             //id
	UserId         string `json:"userId"`         //用户userid
	FileName       string `json:"fileName"`       //文件名字
	Ptime          int64  `json:"ptime"`          //时间
	FileCid        string `json:"fileCid"`        //文件cid
	FileSize       int64  `json:"fileSize"`       //文件大小
	DownPath       string `json:"downPath"`       //下载路径
	FileType       int64  `json:"fileType"`       //文件类型
	TransferType   int64  `json:"transferType"`   //传输类型
	UploadParentId string `json:"uploadParentId"` //下载父id
	UploadFileId   string `json:"uploadFileId"`   //下载文件的id
}

//article

type Article struct {
	Id             string `json:"id"`             //id
	UserId         string `json:"userId"`         //用户id
	Accesstory     string `json:"accesstory"`     //附件
	AccesstoryType int64  `json:"accesstoryType"` //附件类型
	Text           string `json:"text"`           //文本信息
	Tag            string `json:"tag"`            //标签
	Ptime          int64  `json:"ptime"`          //创建时间
	PlayNum        int64  `json:"playNum"`        //播放数量
	ShareNum       int64  `json:"shareNum"`       //分享次数
	Title          string `json:"title"`          //标题
	Thumbnail      string `json:"thumbnail"`      //缩略图
	FileName       string `json:"fileName"`       //文件名字
	FileSize       string `json:"fileSize"`       //文件大小
	LikeNum        int64  `json:"likeNum"`        //数量
	IsLike         int64  `json:"isLike"`         //是否点赞
}
type ArticleAboutMeResp struct {
	Id             string `json:"id"`             //id
	UserId         string `json:"userId"`         //用户id
	Accesstory     string `json:"accesstory"`     //附件
	AccesstoryType int64  `json:"accesstoryType"` //附件类型
	Text           string `json:"text"`           //文本信息
	Tag            string `json:"tag"`            //标签
	Ptime          int64  `json:"ptime"`          //创建时间
	PlayNum        int64  `json:"playNum"`        //播放数量
	ShareNum       int64  `json:"shareNum"`       //分享次数
	Title          string `json:"title"`          //标题
	Thumbnail      string `json:"thumbnail"`      //缩略图
	FileName       string `json:"fileName"`       //文件名字
	FileSize       string `json:"fileSize"`       //文件大小
	Count          int64  `json:"count"`          //数量
	IsLike         int64  `json:"isLike"`         //是否点赞
	PeerId         string `json:"peerId"`         //节点id
	Name           string `json:"name"`           //用户名字
	Phone          string `json:"phone"`          //手机号
	Sex            int64  `json:"sex"`            //性别 0 未知  1 男  2 女
	NickName       string `json:"nickName"`       //昵称
	Img            string `json:"img"`            //头像
	LikeNum        int64  `json:"likeNum"`
}

//article like
type ArticleLike struct {
	Id        string `json:"id"`        //id
	UserId    string `json:"userId"`    //用户id
	ArticleId string `json:"articleId"` //文章id
	IsLike    int64  `json:"isLike"`    //是否点赞
}

// chat_msg
type ChatMsg struct {
	Id          string `json:"id"`          //id
	ContentType int64  `json:"contentType"` //内容类型
	Content     string `json:"content"`     //内容
	FromId      string `json:"fromId"`      //发送者
	ToId        string `json:"toId"`        //接收者
	Ptime       int64  `json:"ptime"`       //创建时间
	IsWithdraw  int64  `json:"isWithdraw"`  //是否撤回         0 未撤回  1  撤回
	IsRead      int64  `json:"isRead"`      //是否已读
	RecordId    string `json:"recordId"`    //房间id
	SendState   int64  `json:"sendState"`   //发送状态  0 待发送或已发送  1 发送成功  -1 发送失败
	SendFail    string `json:"sendFail"`    // 发送失败原因
}

// chat_record

type ChatRecord struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Img     string `json:"img"`
	FromId  string `json:"fromId"`
	Ptime   int64  `json:"ptime"`
	LastMsg string `json:"lastMsg"`
	Toid    string `json:"toId"`
}

type CopyParams struct {
	Pid      string
	CopyFile []File
}
type MoveParams struct {
	Pid      string
	MoveFile []File
}

//delete one file

type DeleteOneParams struct {
	DropFile []File
}

//delete many file

type DeleteManyParams struct {
	DropFile []File
}

//delete one file

type DeleteOneDirParams struct {
	DropFile []File
}

//delete one file

type DeleteManyDirParams struct {
	DropFile []File
}

//delete
type DeleteParams struct {
	DropFile []File
}
