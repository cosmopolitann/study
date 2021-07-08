package jwt

import (
	"encoding/base64"
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
	UserId   string
	PeerId   string `json:"peerId"`   //节点id
	Name     string `json:"name"`     //用户名字
	Phone    string `json:"phone"`    //手机号
	Sex      int64  `json:"sex"`      //性别 0 未知  1 男  2 女
	NickName string `json:"nickName"` //昵称
	Ptime    int64  `json:"-"`        //时间
	Utime    int64  `json:"-"`        //更新时间
	Img      string `json:"img"`      //头像
	jwt.StandardClaims
}

const (
	tokenStr = "adsfa#^$%#$fgrf" //houxu fengzhuang dao nacos
)

func GenerateToken(id, peerid, name, phone, nick, img string, sex, pt, ut int64, expireDuration int64) (string, error) {
	// 将 uid，用户角色， 过期时间作为数据写入 token 中

	calim := LoginClaims{
		UserId:         id,
		PeerId:         peerid,
		Name:           name,
		Phone:          phone,
		Sex:            sex,
		NickName:       nick,
		Ptime:          pt,
		Utime:          ut,
		Img:            img,
		StandardClaims: jwt.StandardClaims{},
	}
	if expireDuration != -1 {
		calim.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireDuration,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)
	//base 64
	// strbytes := []byte(tokenStr)
	// encoded := base64.StdEncoding.EncodeToString(strbytes)

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
	token, err := GenerateToken("123", "perr", "name", "phone", "nick", "img", 1, 2, 3, 60)

	if err != nil {
		t.Log("jwt is failed.")
	}
	//

	t.Log("Token = ", token)

}
