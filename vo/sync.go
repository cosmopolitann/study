package vo

//同步 文章

type SyncArticleAddParams struct {
	Method string           `json:"method"` //同步方法
	Data   ArticleAddParams `json:"data"`   //同步数据
}

// 同步 用户

type SyncUserParams struct {
	Method string           `json:"method"` //同步方法
	Data   ArticleAddParams `json:"data"`   //同步数据
}
type SyncParams struct {
	Method string           `json:"type"` //同步方法
	Data   ArticleAddParams `json:"data"` //同步数据
}

type SyncMsgParams struct {
	Method string      `json:"type"` //同步方法
	Data   interface{} `json:"data"` //同步数据
	FromId string      `json:"from"` //发送者 PeerId
}

//用户

type SyncRecieveUsesrParams struct {
	Method string      `json:"type"` //同步方法
	Data   SyncSysUser `json:"data"` //同步数据
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
	Method string               `json:"type"` //同步方法
	Data   ArticlePlayAddParams `json:"data"` //同步数据
}

//分享次数

type SyncRecieveShareAddParams struct {
	Method string               `json:"type"` //同步方法
	Data   ArticlePlayAddParams `json:"data"` //同步数据
}

// article

type SyncRecieveArticleParams struct {
	Method string           `json:"type"` //同步方法
	Data   ArticleAddParams `json:"data"` //同步数据
}

// 点赞
type SyncRecieveLikeParams struct {
	Method string            `json:"type"` //同步方法
	Data   ArticleLikeParams `json:"data"` //同步数据
}

// 用户更新
type SyncRecieveUserUpdateParams struct {
	Method string      `json:"type"` //同步方法
	Data   SyncSysUser `json:"data"` //同步数据
}

// 取消点赞
type SyncRecieveCancelLikeParams struct {
	Method string            `json:"type"` //同步方法
	Data   ArticleLikeParams `json:"data"` //同步数据
}

type ArticleLikeParams struct {
	Id        string `json:"id"`        //id
	UserId    string `json:"userId"`    //用户id
	ArticleId string `json:"articleId"` //文章id
	IsLike    int64  `json:"isLike"`    //是否点赞
}

//cid
type AddCid struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}
type JwtSysUser struct {
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
