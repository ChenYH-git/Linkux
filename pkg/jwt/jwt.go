package jwt

import (
	"errors"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
)

var mySec = []byte("giao!")

type MyClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

type MyAdClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成Token
func GenToken(userID string, username string) (string, error) {
	c := MyClaims{
		userID,
		username, //自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "linkux",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的字符串Token
	return token.SignedString(mySec)
}

// GenAdToken 生成管理员Token
func GenAdToken(username string) (string, error) {
	c := MyAdClaims{
		username, //自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour).Unix(),
			Issuer:    "linkux",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的字符串Token
	return token.SignedString(mySec)
}

// ParseAdToken 解析管理员Token
func ParseAdToken(tokenString string) (*MyAdClaims, error) {
	mc := new(MyAdClaims)

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
