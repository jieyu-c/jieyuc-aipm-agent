package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jieyuc.cn/jieyuc-aipm-agent/internal/model/users"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel users.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UsersModel: users.NewUsersModel(sqlx.NewSqlConn("postgres", c.DataSource), c.Cache),
	}
}
