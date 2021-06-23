package vo

//前端传参过来的参数  查询列表信息

type CloudFindListParams struct {
	//Name string
	Token    string `json:"token"`
	ParentId string `json:"parentId"`
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
}

//folder
type CloudAddFolderParams struct {
	Id        string `json:"id"`
	FileName  string `json:"fileName"`
	ParentId  string `json:"parentId"`
	Token     string `json:"token"`
	Thumbnail string `json:"thumbnail"` //缩略图
}

//file list
type CloudFileListParams struct {
	ParentId string `json:"parentId"`
}

//folder list
type CloudFolderListParams struct {
	Token    string `json:"token"`
	ParentId string `json:"parentId"`
}

//删除文件和文件夹

type CloudDeleteParams struct {
	Ids   []string `json:"ids"`
	Token string   `json:"token"`
}

//transferadd

type TransferAdd struct {
	Id             string `json:"id"`
	FileName       string `json:"fileName"`
	FileCid        string `json:"fileCid"`
	FileSize       int64  `json:"fileSize"`
	FilePath       string `json:"filePath"`
	FileType       int64  `json:"fileType"`
	TransferType   int64  `json:"transferType"`
	UploadParentId string `json:"uploadParentId"`
	UploadFileId   string `json:"uploadFileId"`
	Token          string `json:"token"`

	//1 上传    2 下载
}

//文件进行分类

type FileCategoryParams struct {
	Token    string `json:"token"`
	FileType int64  `json:"fileType"`
	Order    string `json:"order"`
}

//删除传输列表

type TransferDelParams struct {
	Token string   `json:"token"`
	Ids   []string `json:"ids"`
}

//获取 传输 记录

type TransferListParams struct {
	Token string `json:"token"`
}

// 复制文件

type FileParam struct {
	Id       string `json:"id"`
	UserId   string `json:"userId"`
	FileName string `json:"fileName"`
	ParentId string `json:"parentId"`
	FileCid  string `json:"fileCid"`
	FileSize int64  `json:"fileSize"`
	FileType int64  `json:"fileType"`
	IsFolder int64  `json:"isFolder"`
	Ptime    int64  `json:"-`
}
type CopyFileParams struct {
	Token    string   `json:"token"`
	ParentId string   `json:"parentId"`
	Ids      []string `json:"ids"`
}

//移动文件

type MoveFileParams struct {
	Token    string   `json:"token"`
	ParentId string   `json:"parentId"`
	Ids      []string `json:"ids"`
}

// 查询文件

type SearchFileParams struct {
	Token   string `json:"token"`
	Content string `json:"content"`
	Order   string `json:"order"`
}

// 文章查询

type ArticleSearchParams struct {
	PageSize int64 `json:"pageSize"` // 一次多少条
	PageNum  int64 `json:"pageNum"`  // 第几页
	//Token    string   `json:"token"`
	Title string `json:"title"`
}
