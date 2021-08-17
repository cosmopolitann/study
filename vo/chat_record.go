package vo

//消息记录新增 (类似房间)

type ChatRecoredAddParams struct {
	Id       string `json:"id"`
	Name     string `json:"name"`     //       1 文本  2 表情 3 图片 4 文件
	Img      string `json:"img"`      // require    消息头像
	FromId   string `json:"fromId"`   //require     创建者
	LastMsg  string `json:"lastMsg"`  //require     最后的消息
	ToId     string `json:"toId"`     //require         0 未撤回  1  撤回
	IsActive int64  `json:"isActive"` //require         0 用传进来的  1  生成新的
	Token    string `json:"token"`
}

//获取消息记录列表

type ChatRecordListParams struct {
	FromId     string `json:"fromId"`
	Token      string `json:"token"`
	CustomerId string `json:"customerId"`
	Keyword    string `json:"keyword"`
}

// 获取消息记录列表

type ChatRecordDelParams struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

//
type ChatRecordInfo struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Img      string `json:"img"`
	FromId   string `json:"fromId"`
	Ptime    int64  `json:"ptime"`
	LastMsg  string `json:"lastMsg"`
	Toid     string `json:"toId"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	PeerId   string `json:"peerId"`
	NickName string `json:"nickName"`
	Sex      int64  `json:"sex"`
}
type ChatRecordRespListParams struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Img          string `json:"img"`
	LastMsg      string `json:"lastMsg"`
	Ptime        int64  `json:"ptime"`
	UnreadMsgNum int64  `json:"unreadMsgNum"`

	FromId       string `json:"fromId"`
	FromName     string `json:"fromName"`
	FromImg      string `json:"fromImg"`
	FromPhone    string `json:"fromPhone"`
	FromPeerId   string `json:"fromPeerId"`
	FromNickName string `json:"fromNickName"`
	FromSex      int64  `json:"fromSex"`

	ToId       string `json:"toId"`
	ToName     string `json:"toName"`
	ToImg      string `json:"toImg"`
	ToPhone    string `json:"toPhone"`
	ToPeerId   string `json:"toPeerId"`
	ToNickName string `json:"toNickName"`
	ToSex      int64  `json:"toSex"`
}

type ChatRenameRecordParams struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ChatRecored struct {
	Id      string `json:"id"`
	Name    string `json:"name"`    //       1 文本  2 表情 3 图片 4 文件
	FromId  string `json:"fromId"`  //require     创建者
	LastMsg string `json:"lastMsg"` //require     最后的消息
	ToId    string `json:"toId"`    //require         0 未撤回  1  撤回
	Ptime   int64  `json:"ptime"`   // require    时间
}
