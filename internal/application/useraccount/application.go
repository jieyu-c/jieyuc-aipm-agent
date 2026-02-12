package useraccount

import (
	"context"

	"jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
)

// Service 用户账户应用服务，编排领域与仓储，供 RPC/HTTP/定时任务等调用
type Service struct {
	repo useraccount.UserRepository
	dom  useraccount.DomainService
}

// NewService 构造应用服务
func NewService(repo useraccount.UserRepository) *Service {
	return &Service{repo: repo, dom: useraccount.DomainService{}}
}

// GetUserDetail 获取用户详情（仅后端/内部调用）
func (s *Service) GetUserDetail(ctx context.Context, userId string) (*useraccount.User, error) {
	return s.repo.FindByUserId(ctx, userId)
}
