package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Log   logx.LogConf
	DB    DB `json:"DB"`
	Cache cache.CacheConf
}

type DB struct {
	Member string `json:"member"`
}
