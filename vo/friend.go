package vo

type FriendUpdateNicknameParams struct {
	Token    string `json:"token"`    //token
	FriendId string `json:"friendId"` //friendId
	Nickname string `json:"nickname"` // nickname
}

type FriendCheckOnlineParams struct {
	FriendIds []string `json:"friendIds"` //friendIds
	Token     string   `json:"token"`     //token
}

type FriendSwapOnlineParams struct {
	FromId string `json:"fromId"` // fromid
	ToId   string `json:"toId"`   // toid
}
