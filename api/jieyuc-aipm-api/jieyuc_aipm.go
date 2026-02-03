// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"

	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/bootstrap"
)

var configFile = flag.String("f", "etc/jieyuc_aipm.yaml", "the config file")

func main() {
	flag.Parse()

	server, err := bootstrap.NewAPIService(*configFile)
	if err != nil {
		panic(err)
	}
	defer server.Stop()

	server.Start()
}
