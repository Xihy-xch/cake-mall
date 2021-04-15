package svc

import (
	"cake-mall/api/internal/config"
	"cake-mall/api/internal/public_perform/wechat"
	"cake-mall/rpc/member-data/memberclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	MemberClient memberclient.Member
	WechatClient *wechat.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		MemberClient: memberclient.NewMember(zrpc.MustNewClient(c.Member)),
		WechatClient: wechat.NewClient(c),
	}
}
