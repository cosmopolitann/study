package vo

//前端传参过来的参数  查询列表信息

type CloudFindListParams struct {
	//Name string
	Token    string `json:"token"`    //token
	ParentId string `json:"parentId"` //父id
}

//file
type CloudAddFileParams struct {
	Id        string `json:"id"`
	FileName  string `json:"fileName"`  //文件名字
	ParentId  string `json:"parentId"`  //父id
	FileCid   string `json:"fileCid"`   //文件cid
	FileSize  int64  `json:"fileSize"`  //文件大小
	FileType  int64  `json:"fileType"`  //文件类型
	Token     string `json:"token"`     //token
	Thumbnail string `json:"thumbnail"` //缩略图
	Duration  int64  `json:"duration"`  //时长
	Width     string `json:"width"`     //宽
	Height    string `json:"height"`    //高
}

//folder
type CloudAddFolderParams struct {
	Id        string `json:"id"`        //id
	FileName  string `json:"fileName"`  //文件名字
	ParentId  string `json:"parentId"`  //父id
	Token     string `json:"token"`     //token
	Thumbnail string `json:"thumbnail"` //缩略图
}

//file list
type CloudFileListParams struct {
	ParentId string `json:"parentId"` //父id
}

//folder list
type CloudFolderListParams struct {
	Token    string `json:"token"`    //token
	ParentId string `json:"parentId"` //父id
}

//删除文件和文件夹

type CloudDeleteParams struct {
	Ids   []string `json:"ids"`   //ids
	Token string   `json:"token"` //token
}

//transferadd

type TransferAdd struct {
	Id             string `json:"id"`             //id
	FileName       string `json:"fileName"`       //文件名字
	FileCid        string `json:"fileCid"`        //文件cid
	FileSize       int64  `json:"fileSize"`       //文件大小
	FilePath       string `json:"filePath"`       //文件路径
	FileType       int64  `json:"fileType"`       //文件类型
	TransferType   int64  `json:"transferType"`   //传输类型  0 上传 1 下载
	UploadParentId string `json:"uploadParentId"` //上传父id
	UploadFileId   string `json:"uploadFileId"`   //上传文件id
	Token          string `json:"token"`          //token

	//1 上传    2 下载
}

//文件进行分类

type FileCategoryParams struct {
	Token    string `json:"token"`    //token
	FileType int64  `json:"fileType"` //文件类型
	Order    string `json:"order"`    //排序
}

//删除传输列表

type TransferDelParams struct {
	Token string   `json:"token"` //token
	Ids   []string `json:"ids"`   //ids
}

//获取 传输 记录

type TransferListParams struct {
	Token string `json:"token"` //token
}

// 复制文件

type FileParam struct {
	Id       string `json:"id"`       //id
	UserId   string `json:"userId"`   //用户id
	FileName string `json:"fileName"` //文件名字
	ParentId string `json:"parentId"` //父id
	FileCid  string `json:"fileCid"`  //文件cid
	FileSize int64  `json:"fileSize"` //文件大小
	FileType int64  `json:"fileType"` //文件类型
	IsFolder int64  `json:"isFolder"` //是否是文件夹
	Ptime    int64  `json:"-`         //创建时间
}
type CopyFileParams struct {
	Token    string   `json:"token"`    //token
	ParentId string   `json:"parentId"` //父id
	Ids      []string `json:"ids"`      //ids
}

//移动文件

type MoveFileParams struct {
	Token    string   `json:"token"`    //token
	ParentId string   `json:"parentId"` //父id
	Ids      []string `json:"ids"`      //ids
}

// 查询文件

type SearchFileParams struct {
	Token   string `json:"token"`   //token
	Content string `json:"content"` //内容
	Order   string `json:"order"`   //排序类型
}

// 文章查询

type ArticleSearchParams struct {
	PageSize int64 `json:"pageSize"` // 一次多少条
	PageNum  int64 `json:"pageNum"`  // 第几页
	//Token    string   `json:"token"`
	Title string `json:"title"`
}

//
type ArticleSearchCategoryParams struct {
	PageSize int64 `json:"pageSize"` // 一次多少条
	PageNum  int64 `json:"pageNum"`  // 第几页
	//Token    string   `json:"token"`
	AccesstoryType int64 `json:"accesstoryType"`
}

// sync query all data

type QueryAllData struct {
	Token string `json:"token"` //token
}
type DatabaseMigrationParams struct {
	Token string `json:"token"` //token
}
