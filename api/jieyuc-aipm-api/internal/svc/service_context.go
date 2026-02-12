// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/config"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/middleware"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/user-account/useraccount"
)

type ServiceContext struct {
	Config         config.Config
	UserAccountRpc useraccount.UserAccount
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserAccountRpc: useraccount.NewUserAccount(zrpc.MustNewClient(c.UserAccountRpcConf)),
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
