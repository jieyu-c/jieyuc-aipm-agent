package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	useraccountapp "jieyuc.cn/jieyuc-aipm-agent/internal/application/useraccount"
	useraccountdomain "jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/internal/infrastructure/repository/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/internal/model/users"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/user-account/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	UserRepository useraccountdomain.UserRepository
	UserAccountApp *useraccountapp.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	usersModel := users.NewUsersModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache)
	repo := useraccount.NewPostgresUserRepository(usersModel)
	return &ServiceContext{
		Config:         c,
		UserRepository: repo,
		UserAccountApp: useraccountapp.NewService(repo),
	}
}
