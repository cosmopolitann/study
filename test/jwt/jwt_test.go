package jwt

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

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
	Id       string `json:"id"`
	PeerId   string `json:"peerId"`   //节点id
	Name     string `json:"name"`     //用户名字
	Phone    string `json:"phone"`    //手机号
	Sex      int64  `json:"sex"`      //性别 0 未知  1 男  2 女
	NickName string `json:"nickname"` //昵称
	Img      string `json:"img"`      //头像
	Ptime    int64  `json:"ptime"`    //时间
	Utime    int64  `json:"utime"`    //更新时间
	Role     string `json:"role"`
	jwt.StandardClaims
}

const (
	tokenStr = "adsfa#^$%#$fgrf" //houxu fengzhuang dao nacos
)

func GenerateToken(id, peerId, name, phone, nickname, img, role string, sex, ptime, utime int64, expireDuration int64) (string, error) {
	// 将 uid，用户角色， 过期时间作为数据写入 token 中

	calim := LoginClaims{
		Id:             id,
		PeerId:         peerId,
		Name:           name,
		Phone:          phone,
		Sex:            sex,
		NickName:       nickname,
		Img:            img,
		Ptime:          ptime,
		Utime:          utime,
		Role:           role,
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
	fmt.Println("开始")
	// 414207114215428096', '', '','', 0, 1627444008, 1627444008, '人工客服1', ''
	//                  id, peerId, name, phone, nickname, img string, sex, ptime, utime int64
	token, err := GenerateToken("414207114215428090", "", "测试", "", "测试", "", "1", 0, 1627444008, 1627444008, 60*60*60)

	if err != nil {
		t.Log("jwt is failed.")
	}
	t.Log("Token = ", token)
	fmt.Println("token:=", token)
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI0MTY5ODQ1NDUwNjIwMzEzNjAiLCJleHAiOjE2MjYzNTUxMTl9.Ko9C6ojPzShQ3BSP_ASa602EUjD27trRO_11zaV4hCY
