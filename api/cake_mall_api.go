package main

import (
	"flag"
	"fmt"

	"cake-mall/api/internal/config"
	"cake-mall/api/internal/handler"
	"cake-mall/api/internal/svc"
	server2 "cake-mall/api/middlerware/http/server"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "/home/xch/go/src/cake-mall/api/etc/cake_mall_api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	server.Use(server2.LogServerMiddleware)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
