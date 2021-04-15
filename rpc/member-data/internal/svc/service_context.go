package svc

import (
	"cake-mall/rpc/member-data/internal/config"
	"cake-mall/rpc/member-data/model"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	UserPwdModel    model.UserPwdModel
	WechatUserModel model.WechatUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserPwdModel:    model.NewUserPwdModel(sqlx.NewMysql(c.DB.Member), c.Cache),
		WechatUserModel: model.NewWechatUserModel(sqlx.NewMysql(c.DB.Member), c.Cache),
	}
}
