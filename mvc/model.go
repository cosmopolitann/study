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
	Id             string `json:"id"`
	UserId         string `json:"userId"`
	Accesstory     string `json:"accesstory"`
	AccesstoryType int64  `json:"accesstoryType"`
	Text           string `json:"text"`
	Tag            string `json:"tag"`
	Ptime          int64  `json:"ptime"`
	PlayNum        int64  `json:"playNum"`
	ShareNum       int64  `json:"shareNum"`
	Title          string `json:"title"`
	Thumbnail      string `json:"thumbnail"`
	FileName       string `json:"fileName"`
	FileSize       string `json:"fileSize"`
	Count          int64  `json:"count"`
}

//article like
type ArticleLike struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	ArticleId string `json:"articleId"`
	IsLike    int64  `json:"isLike"`
}

// chat_msg
type ChatMsg struct {
	Id          string `json:"id"`
	ContentType int64  `json:"contentType"`
	Content     string `json:"content"`
	FromId      string `json:"fromId"`
	ToId        string `json:"toId"`
	Ptime       int64  `json:"ptime"`
	IsWithdraw  int64  `json:"isWithdraw"` //require     coment 是否撤回         0 未撤回  1  撤回
	IsRead      int64  `json:"isRead"`
	RecordId    string `json:"recordId"`
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
	//sql := fmt.Sprintf("INSERT INTO article (id,  user_id      ,accesstory,    accesstory_type,    text,      tag,     ptime,   play_num,share_num,title,     thumbnail,file_name,file_size) values 
	//                                        ('%s','%s',        '%s'           ,%d,                  '%s',  '    %s'    ,%d,          %d     ,%d,     '%s',   '     %s','    %s',     %d)\n", 
	//                                         sid, art.UserId,  art.Accesstory, art.AccesstoryType, art.Text, art.Tag, t,           0,       0,    art.Title, art.Thumbnail, art.FileName, art.FileSize)
	