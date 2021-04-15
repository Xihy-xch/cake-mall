package svc

import (
	"cake-mall/api/internal/config"
	"cake-mall/api/internal/public_perform/wechat"
)

type ServiceContext struct {
	Config       config.Config
	WechatClient *wechat.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		WechatClient: wechat.NewClient(c),
	}
}
