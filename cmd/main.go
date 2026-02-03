package main

import (
	"flag"
	"fmt"
	"log"

	apiBootstrap "jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/bootstrap"
	_ "jieyuc.cn/jieyuc-aipm-agent/internal/model"
	rpcBootstrap "jieyuc.cn/jieyuc-aipm-agent/service/user-account/bootstrap"

	"github.com/zeromicro/go-zero/core/service"
)

var (
	apiConfigFile = flag.String("api-f", "api/jieyuc-aipm-api/etc/jieyuc_aipm.yaml", "the api config file")
	rpcConfigFile = flag.String("rpc-f", "service/user-account/etc/useraccount-local.yaml", "the rpc config file")
)

func main() {
	flag.Parse()

	group := service.NewServiceGroup()
	defer group.Stop()

	// Initialize API Service
	apiServer, err := apiBootstrap.NewAPIService(*apiConfigFile)
	if err != nil {
		log.Fatalf("Failed to create API service: %v", err)
	}
	group.Add(apiServer)

	// Initialize RPC Service
	rpcServer, err := rpcBootstrap.NewRPCService(*rpcConfigFile)
	if err != nil {
		log.Fatalf("Failed to create RPC service: %v", err)
	}
	group.Add(rpcServer)

	fmt.Println("Starting services...")
	group.Start()
}
