package jwt

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/cosmopolitann/clouddb/vo"
	"github.com/dgrijalva/jwt-go"
)

/*
token, err := utils.GenerateToken(
		user.StariverUserId,
		user.NickName,
		user.Mobile,
		user.Email,
		"", 30*24*60*60)
*/
type LoginClaims struct {
	UserId   string
	PeerId   string `json:"peerId"`   //节点id
	Name     string `json:"name"`     //用户名字
	Phone    string `json:"phone"`    //手机号
	Sex      int64  `json:"sex"`      //性别 0 未知  1 男  2 女
	NickName string `json:"nickName"` //昵称
	Img      string `json:"img"`      //头像
	jwt.StandardClaims
}

const (
	tokenStr = "adsfa#^$%#$fgrf" //houxu fengzhuang dao nacos
)

func GenerateToken(user vo.RespSysUser, expireDuration int64) (string, error) {
	// 将 uid，用户角色， 过期时间作为数据写入 token 中

	calim := LoginClaims{
		UserId:         user.Id,
		PeerId:         user.PeerId,
		Name:           user.Name,
		Phone:          user.Phone,
		Sex:            user.Sex,
		NickName:       user.NickName,
		Img:            user.Img,
		StandardClaims: jwt.StandardClaims{},
	}
	if expireDuration != -1 {
		calim.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireDuration,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)
	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)

	return token.SignedString(strBase)
}

//func ParseToken(strGen string) (*jwt.Token, error) {
//	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)
//	return jwt.Parse(strGen, func(*jwt.Token) (interface{}, error) {
//		return strBase, nil
//	})
//}
func TestJwt(t *testing.T) {
	//token,err:=GenerateToken("10001",30*24*60*60)
	info := vo.RespSysUser{
		Id:       "123",
		Name:     "lily",
		PeerId:   "Qm123",
		Phone:    "1233",
		Sex:      0,
		NickName: "lili",
		Img:      "llllll",
	}
	token, err := GenerateToken(info, 60*60*60)

	if err != nil {
		t.Log("jwt is failed.")
	}
	t.Log("Token = ", token)

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTY5ODQ1NDUwNjIwMzEzNjAiLCJleHAiOjE2MjYzNTUxMTl9.Ko9C6ojPzShQ3BSP_ASa602EUjD27trRO_11zaV4hCY
