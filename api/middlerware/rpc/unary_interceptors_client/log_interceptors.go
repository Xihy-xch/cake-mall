package unary_interceptors_client

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/timex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func LogInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	startTime := timex.Now()
	var addr string
	client, ok := peer.FromContext(ctx)
	if ok {
		addr = client.Addr.String()
	}

	logx.WithContext(ctx).Infof("发送请求[RPC]：[addr: %s - method: %s - req: %s]", addr, method, req)

	err := invoker(ctx, method, req, reply, cc, opts...)

	logx.WithContext(ctx).WithDuration(timex.Since(startTime)).Infof("接收响应[RPC]：[addr: %s - method: %s - req: %s]", addr, method, reply)
	return err
}
