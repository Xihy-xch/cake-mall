package main

import (
	"flag"
	"fmt"

	"cake-mall/rpc/member-data/internal/config"
	"cake-mall/rpc/member-data/internal/server"
	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "/home/xch/go/src/cake-mall/rpc/member-data/etc/member_data.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewMemberServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		reflection.Register(grpcServer)
		member.RegisterMemberServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
