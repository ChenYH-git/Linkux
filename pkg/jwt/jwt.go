package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

var mySec = []byte("giao!")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成Token
func GenToken(userID int64, username string) (string, error) {
	c := MyClaims{
		userID,
		username, //自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "bluebell",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的字符串Token
	return token.SignedString(mySec)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	mc := new(MyClaims)

	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySec, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, err
	}
	return nil, errors.New("invalid token")
}
