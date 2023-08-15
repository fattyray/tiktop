package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	UserId   string `json:"userid"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

var MySecret = []byte("字节后端小分队")

// 生成token
func CreateToken(userId string, password string) (string, error) {
	claim := MyClaims{
		UserId:   userId,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1) * 8)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "字节后端小分队",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token_str, err := token.SignedString(MySecret)
	return token_str, err
}

func Gettoken(token string) (*MyClaims, error) {
	mytoken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if mytoken != nil {
		if claims, ok := mytoken.Claims.(*MyClaims); ok && mytoken.Valid {
			return claims, nil
		}
	}
	return nil, err
}
