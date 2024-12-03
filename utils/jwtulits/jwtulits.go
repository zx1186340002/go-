package jwtulits

import (
	"fmt"
	"program/config"
	"program/utils/logs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSignKey = []byte(config.JwtSignKey)

type MycustomClaim struct {
	username string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(username string) (string, error) {
	claim := MycustomClaim{ //载荷数据
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Zhouxin",
			Subject:   "example",
		},
	}
	logsa := make(map[string]interface{})
	logsa["Username"] = claim.username
	logsa["ExpiresAt"] = claim.ExpiresAt
	logsa["NotBefore"] = claim.NotBefore
	logsa["Issuer"] = claim.Issuer
	logsa["Subject"] = claim.Subject
	logs.Debug(logsa, "生成消息")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //加密
	fmt.Println(token)
	ss, err := token.SignedString(jwtSignKey) //签名
	fmt.Println(ss)
	return ss, err
}
