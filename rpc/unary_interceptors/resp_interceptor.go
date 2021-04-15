package unary_interceptors

import (
	"context"
	"encoding/json"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/timex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const serverSlowThreshold = time.Millisecond * 500

func UnaryLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		startTime := timex.Now()
		resp, err = handler(ctx, req)
		logDuration(ctx, info.FullMethod, resp, timex.Since(startTime), err)

		return
	}
}

func logDuration(ctx context.Context, method string, resp interface{}, duration time.Duration, respErr error) {
	var addr string
	client, ok := peer.FromContext(ctx)
	if ok {
		addr = client.Addr.String()
	}
	content, err := json.Marshal(resp)

	if err != nil {
		logx.WithContext(ctx).Errorf("%s - %s", addr, err.Error())
	} else if duration > serverSlowThreshold {
		logx.WithContext(ctx).WithDuration(duration).Slowf("发送响应[RPC] [slowcall - addr: %s - method: %s - resp: %s]",
			addr, method, string(content))
	} else {
		logx.WithContext(ctx).WithDuration(duration).Infof("发送响应[RPC]：[addr: %s - method: %s - resp: %s]", addr, method, string(content))
	}
}
