package jwt

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"log"
	"strings"
	"testing"
)

const (
	TOKEN_ERR_NONE    = "0"
	TOKEN_ERR_LEN     = "1"
	TOKEN_ERR_EXPIRED = "2"
)

func TestParseJwt(t *testing.T) {
	// parse

	claim, flag := GetClaim("")
	if flag != TOKEN_ERR_LEN && flag != TOKEN_ERR_EXPIRED {
		t.Log("T--------")
	}
	t.Log("calim is  userId === ", claim)

	t.Log("calim is  userId === ", claim["UserId"])

}
func GetClaim(bareStr string) (jwt.MapClaims, string) {
	bareStr = "Auth eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDAwMSIsImV4cCI6MTYyNTc0ODM0OH0.YyElfrbh_A1jUg5XzPWbF1HPvuYVwCl77h-Eoh7o05s"
	bareArr := strings.Split(bareStr, " ")
	errFlag := TOKEN_ERR_NONE
	if len(bareArr) != 2 {
		errFlag = TOKEN_ERR_LEN
		log.Println(" 错误 ：=", errFlag)
	}
	token, err := ParseToken(bareArr[1])
	log.Println(" token = ", token)
	vl := token.Valid
	log.Println("校验结果 = ", vl)
	if err != nil || token.Claims == nil {
		errFlag = TOKEN_ERR_EXPIRED
		return nil, errFlag
	}
	claim := token.Claims.(jwt.MapClaims)
	return claim, errFlag
}
func ParseToken(strGen string) (*jwt.Token, error) {
	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)
	return jwt.Parse(strGen, func(*jwt.Token) (interface{}, error) {
		return strBase, nil
	})
}
