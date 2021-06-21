package vo

//  user login

type UserLoginParams struct {
	Phone string `json:"phone"`
	Prikey string `json:"prikey"`
}

// 登录返回参数
//
type RespSysUser struct {
	Id       string `json:"id"`
	PeerId   string `json:"peerId"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      int64  `json:"sex"`
	NickName string `json:"nickName"`
	Ptime    int64  `json:"ptime"`
	Utime    int64  `json:"utime"`
	Img      string `json:"img"`
}

type UserLoginRespParams struct {
	Token    string      `json:"token"`
	UserInfo RespSysUser `json:"userInfo"`
}

// User del
type UserDelParams struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}

//  user list

type UserListParams struct {
	//Id string 	`json:"id"`
	Token string `json:"token"`
}
type UserUpdateParams struct {
	Img     string `json:"img"`
	Phone    string `json:"phone"`
	Sex      int64  `json:"sex"`
	NickName string `json:"nickName"`
	Token    string `json:"token"`
}
