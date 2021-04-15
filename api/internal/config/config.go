package config

import (
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Log    logx.LogConf
	Wechat WechatConfig
}

type WechatConfig struct {
	AppID     string `json:"AppID"`
	AppSecret string `json:"AppSecret"`
	Host      string `json:"Host"`
}
