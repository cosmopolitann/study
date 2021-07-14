package jwt

import (
	"encoding/base64"
	"log"
	"strings"
	"time"

	"github.com/cosmopolitann/clouddb/vo"
	"github.com/dgrijalva/jwt-go"
)

const (
	TOKEN_ERR_NONE    = "0"
	TOKEN_ERR_LEN     = "1"
	TOKEN_ERR_EXPIRED = "2"
)

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

func GenerateToken(user vo.RespSysUser, expireDuration int64) (string, error) {
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
