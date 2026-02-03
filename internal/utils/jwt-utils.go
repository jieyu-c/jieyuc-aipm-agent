package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	// JWT 加解密密钥
	DefaultSecretKey = "FE6C380BAE24D12AC952513550173D79"
	// 过期时间，单位秒
	DefaultExpireTime = 12 * time.Hour
)

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体

func getJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetJwtToken(payload string) (string, error) {
	return getJwtToken(DefaultSecretKey, time.Now().Unix(), int64(DefaultExpireTime.Seconds()), payload)
}

// ParseJwtToken 解析 JWT，返回 payload 字符串；token 格式为 "Bearer <token>" 或直接 "<token>"
func ParseJwtToken(tokenString string) (payload string, err error) {
	if tokenString == "" {
		return "", fmt.Errorf("token 不能为空")
	}
	const bearer = "Bearer "
	if len(tokenString) > len(bearer) && tokenString[:len(bearer)] == bearer {
		tokenString = tokenString[len(bearer):]
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(DefaultSecretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("无效的 token")
	}
	p, _ := claims["payload"].(string)
	return p, nil
}
