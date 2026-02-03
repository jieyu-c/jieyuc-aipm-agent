package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	// bcrypt 哈希前缀，用于识别已加密的密码
	bcryptPrefix = "$2"
	// bcrypt 默认成本因子
	bcryptCost = bcrypt.DefaultCost
)

// HashPassword 使用 bcrypt 对密码进行加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// VerifyPassword 验证密码是否正确
// 支持 bcrypt 哈希和明文密码（用于兼容旧数据，新注册用户应使用 bcrypt）
func VerifyPassword(hashedPassword, plainPassword string) bool {
	if hashedPassword == "" || plainPassword == "" {
		return false
	}
	// 已是 bcrypt 哈希，使用 bcrypt 验证
	if strings.HasPrefix(hashedPassword, bcryptPrefix) {
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
		return err == nil
	}
	// 明文存储（兼容旧数据，建议迁移后移除）
	return hashedPassword == plainPassword
}
