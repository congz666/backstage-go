//Package util ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-15 14:37:47
 * @LastEditors: congz
 * @LastEditTime: 2020-09-21 16:01:59
 */
package util

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

//Claims ...
type Claims struct {
	Identification string
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken() (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		os.Getenv("Identification"),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "cmall",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
