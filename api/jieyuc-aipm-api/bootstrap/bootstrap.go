package bootstrap

import (
	"fmt"

	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/config"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/handler"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func NewAPIService(configFile string) (*rest.Server, error) {
	var c config.Config
	conf.MustLoad(configFile, &c)

	server := rest.MustNewServer(c.RestConf)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	return server, nil
}
