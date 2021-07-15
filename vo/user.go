package vo

//  user login

type UserLoginParams struct {
	Phone  string `json:"phone"`
	Prikey string `json:"prikey"`
}

// 登录返回参数

type RespSysUser struct {
	Id       string `json:"id"`       //id
	PeerId   string `json:"peerId"`   //节点id
	Name     string `json:"name"`     //用户名字
	Phone    string `json:"phone"`    //手机号  暂时用
	Sex      int64  `json:"sex"`      //性别 0 未知 1男 2女
	NickName string `json:"nickName"` //昵称
	Ptime    int64  `json:"ptime"`    //创建时间
	Utime    int64  `json:"utime"`    //更新时间
	Img      string `json:"img"`      //图片
}
type RespArticleLike struct {
	Id        string `json:"id"`        //id
	UserId    string `json:"userId"`    //用户
	ArticleId string `json:"articleId"` //文章id
	IsLike    int64  `json:"isLike"`    //点赞
}
type UserLoginRespParams struct {
	Token    string      `json:"token"`    //token
	UserInfo RespSysUser `json:"userInfo"` //用户信息
}

// User del
type UserDelParams struct {
	Id    string `json:"id"`    //id
	Token string `json:"token"` //token
}

//  user list

type UserListParams struct {
	//Id string 	`json:"id"`
	Token string `json:"token"` //token
}
type UserUpdateParams struct {
	Img      string `json:"img"`      //图片
	Phone    string `json:"phone"`    //手机号 暂时用
	Sex      int64  `json:"sex"`      //性别 0 未知 1男  2 女
	NickName string `json:"nickName"` //昵称
	Token    string `json:"token"`    //token
}

// Other User Info.
type OtherUserInfoParams struct {
	UserId string `json:"userid"` //userid
}
