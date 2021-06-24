package vo


//同步 文章

type SyncArticleAddParams struct {
	Method string           `json:"method"`
	Data   ArticleAddParams `json:"data"`
}

// 同步 用户

type SyncUserParams struct {
	Method string           `json:"method"`
	Data   ArticleAddParams `json:"data"`
}
type SyncParams struct {
	Method string           `json:"type"`
	Data   ArticleAddParams `json:"data"`
}

type SyncMsgParams struct {
	Method string           `json:"type"`
	Data   interface{}     `json:"data"`
	FromId string         `json:"from"`
}

//用户

type SyncRecieveUsesrParams struct {
	Method string           `json:"type"`
	Data    SyncSysUser    `json:"data"`
}

type SyncSysUser struct {
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
//播放次数

type SyncRecievePlayParams struct {
	Method string           `json:"type"`
	Data    ArticlePlayAddParams    `json:"data"`
}
//分享次数

type SyncRecieveShareAddParams struct {
	Method string           `json:"type"`
	Data    ArticlePlayAddParams    `json:"data"`
}

// article

type SyncRecieveArticleParams struct {
	Method string            `json:"type"`
	Data    ArticleAddParams    `json:"data"`
}