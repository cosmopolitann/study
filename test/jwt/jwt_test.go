package jwt

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
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
	UserId string
	jwt.StandardClaims
}

const (
	tokenStr = "xiaolongyunpan" //houxu fengzhuang dao nacos
)

func GenerateToken(userId string, expireDuration int64) (string, error) {
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	calim := LoginClaims{
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{},
	}
	if expireDuration != -1 {
		calim.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireDuration,
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)
	//base 64

	
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
	token, err := GenerateToken("10001", 60)

	if err != nil {
		t.Log("jwt is failed.")
	}
	//  

	t.Log("Token = ", token)

}
