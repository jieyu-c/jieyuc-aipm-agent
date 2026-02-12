package useraccount

import (
	"context"
	"errors"
	"fmt"

	useraccountdomain "jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
)

// 统一错误信息，防止通过响应差异枚举用户
const errLoginFailed = "手机号或密码错误"

// LoginResult 登录成功后的结果
type LoginResult struct {
	User  *useraccountdomain.User
	Token string
}

// Login 执行登录：查用户、领域校验、验证凭据、生成 Token
func (s *Service) Login(ctx context.Context, phone, password, verifyCode string) (*LoginResult, error) {
	if phone == "" {
		return nil, fmt.Errorf("手机号不能为空")
	}
	if password == "" && verifyCode == "" {
		return nil, fmt.Errorf("请输入密码或验证码")
	}

	user, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, useraccountdomain.ErrUserNotFound) {
			return nil, fmt.Errorf(errLoginFailed)
		}
		return nil, fmt.Errorf(errLoginFailed)
	}
	if user == nil {
		return nil, fmt.Errorf(errLoginFailed)
	}

	if err := s.dom.CanLogin(user); err != nil {
		return nil, err
	}

	if password != "" {
		if !utils.VerifyPassword(user.Password, password) {
			return nil, fmt.Errorf(errLoginFailed)
		}
	} else {
		// TODO: 验证码登录，需接入验证码服务
		return nil, fmt.Errorf("验证码登录功能暂未开放，请使用密码登录")
	}

	tokenPayload := useraccountdomain.NewTokenPayload(user.UserId, user.Phone)
	payload, err := tokenPayload.ParseToJson()
	if err != nil {
		return nil, fmt.Errorf("登录失败")
	}
	token, err := utils.GetJwtToken(payload)
	if err != nil {
		return nil, fmt.Errorf("登录失败，请稍后重试")
	}

	return &LoginResult{User: user, Token: token}, nil
}
