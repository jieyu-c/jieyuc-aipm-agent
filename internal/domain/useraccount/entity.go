package useraccount

import (
	"encoding/json"
	"errors"
	"time"
)

// 用户状态（与持久化一致）
const (
	UserStatusNormal = 0
	UserStatusBanned = 1
)

// User 用户账户领域实体，与持久化模型解耦，不含 db tag
type User struct {
	UserId    string
	Phone     string
	Username  string
	Password  string
	Nickname  string
	Avatar    string
	Gender    int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TokenPayload struct {
	UserId string `json:"userId"`
	Phone  string `json:"phone"`
}

func NewTokenPayload(userId, phone string) *TokenPayload {
	return &TokenPayload{
		UserId: userId,
		Phone:  phone,
	}
}

// ParseToJson 将TokenPayload转为Json字符串
func (t *TokenPayload) ParseToJson() (string, error) {
	payload, err := json.Marshal(t)
	if err != nil {
		return "", errors.New("转JSON字符串失败")
	}
	return string(payload), nil
}

func Json2TokenPayload(payload string) (*TokenPayload, error) {
	var payloadObj TokenPayload
	err := json.Unmarshal([]byte(payload), &payloadObj)
	if err != nil {
		return nil, err
	}
	return &payloadObj, nil
}
