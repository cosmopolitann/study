package jwt

import (
	"encoding/base64"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TOKEN_ERR_NONE    = "0"
	TOKEN_ERR_LEN     = "1"
	TOKEN_ERR_EXPIRED = "2"
)

type LoginClaims struct {
	UserId string
	jwt.StandardClaims
}

const (
	tokenStr = "adsfa#^$%#$fgrf" //houxu fengzhuang dao nacos
	// tokenStr = "xiaolongyunpan" //houxu fengzhuang dao nacos

	// xiaolongyunpan
)

func GenerateToken(userId string, expireDuration int64) (string, error) {
	calim := LoginClaims{
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{},
	}
	if expireDuration != -1 {
		calim.StandardClaims = jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + expireDuration,
		}
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, calim)

	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)
	return token.SignedString(strBase)
}

func JwtVeriyToken(token string) (t jwt.MapClaims, is bool) {
	token = "Auth " + token
	claim, flag, b := GetClaim(token)

	if flag != TOKEN_ERR_LEN && flag != TOKEN_ERR_EXPIRED {

	}
	if b == false {
		return claim, false
	} else {
		return claim, true
	}
}

func GetClaim(bareStr string) (jwt.MapClaims, string, bool) {
	//bareStr = "Auth eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDAwMSIsImV4cCI6MTYyMzIyNjUxNX0.y1a0_t3fsgohGTxOenA7Lpl5PHll9diyDfwPCPdYxdA"
	bareArr := strings.Split(bareStr, " ")
	errFlag := TOKEN_ERR_NONE
	if len(bareArr) != 2 {
		errFlag = TOKEN_ERR_LEN
		log.Println(" 错误 ：=", errFlag)
	}
	log.Println(" 获取 bareArr= ", bareArr)

	token, err := ParseToken(bareArr[1])
	log.Println(" token 的值= ", token)

	if err != nil || token.Claims == nil {
		errFlag = TOKEN_ERR_EXPIRED
		return nil, errFlag, false
	}

	vl := token.Valid

	log.Println("校验结果 = ", vl)

	claim := token.Claims.(jwt.MapClaims)

	if vl == false {

		return claim, errFlag, false
	}
	return claim, errFlag, true
}
func ParseToken(strGen string) (*jwt.Token, error) {
	strBase, _ := base64.URLEncoding.DecodeString(tokenStr)
	return jwt.Parse(strGen, func(*jwt.Token) (interface{}, error) {
		return strBase, nil
	})
}
