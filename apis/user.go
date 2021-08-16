package apis

import (
	"encoding/json"
	"fmt"

	"github.com/cosmopolitann/clouddb/sugar"
)

const API_PREFIX = "http://cloudapi.stariverpan.com"

func GetUserInfo(authToken, userId string) (uModel UserModel, err error) {

	sugar.Log.Info("GetUserInfo Req --->", authToken, userId)

	apiUrl := API_PREFIX + "/cmsprovider/v1/user/query"
	header := map[string][]string{
		"Authorization": {"Bearer " + authToken},
	}
	data := map[string]interface{}{
		"userId": userId,
	}
	res, err := PostJson(apiUrl, header, data)
	if err != nil {
		err = fmt.Errorf("utils.PostJson err: %w", err)
		return
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		err = fmt.Errorf("json.Marshal err: %w", err)
		return
	}

	err = json.Unmarshal(resBytes, &uModel)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal err: %w", err)
		return
	}

	sugar.Log.Info("GetUserInfo Res <---", res)

	return
}

type UserModel struct {
	Token      string `json:"token"`      // token
	Id         string `json:"id"`         // 用户ID
	PeerId     string `json:"peerId"`     // 用户Peerid
	Name       string `json:"name"`       // 用户名
	Phone      string `json:"phone"`      // 手机号
	Sex        int64  `json:"sex"`        // 性别
	Ptime      int64  `json:"ptime"`      // 创建时间
	Utime      int64  `json:"utime"`      // 修改时间
	Nickname   string `json:"nickname"`   // 昵称
	Img        string `json:"img"`        // 头像
	LikeNum    int64  `json:"likeNum"`    // 点赞数
	ArticleNum int64  `json:"articleNum"` // 发布数
	Role       string `json:"role"`       // 角色

}
