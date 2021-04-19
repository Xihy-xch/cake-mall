package svc

import (
	"cake-mall/rpc/member-data/internal/config"
	"cake-mall/rpc/member-data/model"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	IDGenerator     *IDGenerator
	UserPwdModel    model.UserPwdModel
	WechatUserModel model.WechatUserModel
	UserModel       model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserPwdModel:    model.NewUserPwdModel(sqlx.NewMysql(c.DB.Member), c.Cache),
		WechatUserModel: model.NewWechatUserModel(sqlx.NewMysql(c.DB.Member), c.Cache),
		IDGenerator:     NewIDGenerator(),
	}
}

type IDGenerator struct {
	*snowflake.Snowflake
}

func NewIDGenerator() *IDGenerator {
	sf, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		logx.Error("id生成器初始化错误: ", err)
	}
	return &IDGenerator{sf}
}

func (i *IDGenerator) GetID() int64 {
	return i.NextVal()
}
