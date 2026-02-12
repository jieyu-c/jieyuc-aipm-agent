package main

import (
	"flag"

	"jieyuc.cn/jieyuc-aipm-agent/rpc/user-account/bootstrap"
)

var configFile = flag.String("f", "etc/useraccount.yaml", "the config file")

func main() {
	flag.Parse()

	s, err := bootstrap.NewRPCService(*configFile)
	if err != nil {
		panic(err)
	}
	defer s.Stop()

	s.Start()
}
