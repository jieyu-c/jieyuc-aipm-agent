package useraccount

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	useraccountdomain "jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
)

// RegisterResult 注册结果，业务层不暴露系统错误细节
type RegisterResult struct {
	OK  bool
	Msg string
}

// Register 执行注册：验证码、查重、密码加密、落库
func (s *Service) Register(ctx context.Context, phone, password, verifyCode string) (*RegisterResult, error) {
	if !verifyCodeDev(phone, verifyCode) {
		return &RegisterResult{OK: false, Msg: "验证码错误或已过期"}, nil
	}

	exists, err := s.repo.ExistsByPhone(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}
	if exists {
		return &RegisterResult{OK: false, Msg: "注册失败，请检查输入信息或联系客服"}, nil
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}

	userId := strings.ReplaceAll(uuid.New().String(), "-", "")
	user := &useraccountdomain.User{
		UserId:   userId,
		Phone:    phone,
		Username: userId,
		Password: hashedPassword,
	}
	if err = s.repo.Add(ctx, user); err != nil {
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}

	return &RegisterResult{OK: true, Msg: "注册成功"}, nil
}

// verifyCodeDev 验证码校验（开发用桩，正式需接 Redis/验证码服务）
func verifyCodeDev(phone, code string) bool {
	const devTestCode = "123456"
	return code == devTestCode
}
