package useraccount

import "errors"

var (
	ErrUserNotFound = errors.New("用户不存在")
	ErrUserBanned   = errors.New("账号已被禁用，请联系客服")
)
