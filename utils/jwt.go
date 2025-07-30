package utils

import (
	"Tiktok/config"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// 签名密钥
var jwtSecretKey = []byte(config.SECRETE)

// GenerateToken 生成token
func GenerateToken(userID int64, username string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour).Unix()
	log.Println("expireTime:", expireTime)
	claims := Claims{
		ID:       userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "www.LunBoWang.xyz",
		},
	}
	// 使用用于签名的算法和令牌
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	if token, err := tokenClaims.SignedString(jwtSecretKey); err != nil {
		log.Println("generate token failed")
		fmt.Println(err)
		return "fail"
	} else {
		log.Println("generate token success")
		return token
	}
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
