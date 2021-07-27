package vo

//保存消息

type ChatAddMsgParams struct {
	Id          string `json:"id"`
	ContentType int64  `json:"contentType"` //       1 文本  2 表情 3 图片 4 文件
	Content     string `json:"content"`     // require     coment 消息内容
	FromId      string `json:"fromId"`      //require     coment 发送方id
	ToId        string `json:"toId"`        //require     coment 对方id
	IsWithdraw  int64  `json:"isWithdraw"`  //require     coment 是否撤回         0 未撤回  1  撤回
	IsRead      int64  `json:"isRead"`      // require     coment 是否已读
	RecordId    string `json:"recordId"`    //require     coment 消息记录id
	Token       string `json:"token"`       //token
}

//获取消息分页

type ChatMsgListParams struct {
	PageSize int64  `json:"pageSize"`
	PageNum  int64  `json:"pageNum"`
	RecordId string `json:"recordId"`
	Token    string `json:"token"` //token

}

//
type ArticleListParams struct {
	PageSize int64  `json:"pageSize"`
	PageNum  int64  `json:"pageNum"`
	Token    string `json:"token"` //token
}

type ArticleListUserParams struct {
	PageSize int64  `json:"pageSize"`
	PageNum  int64  `json:"pageNum"`
	UserId   string `json:"userId"` //UserId
}

// 删除消息

type ChatMsgDelParams struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

// 撤回消息

// 删除消息

type ChatMsgWithDrawParams struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
