package bootstrap

import (
	"fmt"

	"jieyuc.cn/jieyuc-aipm-agent/service/pb/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/internal/config"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/internal/server"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewRPCService(configFile string) (*zrpc.RpcServer, error) {
	var c config.Config
	conf.MustLoad(configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_account.RegisterUserAccountServer(grpcServer, server.NewUserAccountServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	return s, nil
}
