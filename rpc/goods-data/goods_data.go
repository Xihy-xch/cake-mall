package main

import (
	"flag"
	"fmt"

	"cake-mall/rpc/goods-data/goods"
	"cake-mall/rpc/goods-data/internal/config"
	"cake-mall/rpc/goods-data/internal/server"
	"cake-mall/rpc/goods-data/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/goodsdata.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewGoodsServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
