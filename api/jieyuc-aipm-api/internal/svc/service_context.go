// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/config"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/useraccount"
)

type ServiceContext struct {
	Config         config.Config
	UserAccountRpc useraccount.UserAccount
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserAccountRpc: useraccount.NewUserAccount(zrpc.MustNewClient(c.UserAccountRpcConf)),
	}
}
