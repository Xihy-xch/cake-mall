package wechat

import (
	"context"
	"fmt"

	"cake-mall/api/internal/config"
	"cake-mall/api/middlerware/http/client"
)

type Client struct {
	AppID     string
	AppSecret string
	*client.HClient
}

func NewClient(config config.Config) *Client {
	return &Client{
		AppID:     config.Wechat.AppID,
		AppSecret: config.Wechat.AppSecret,
		HClient:   client.NewClient(config.Wechat.Host),
	}
}

type AuthCode2SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrorCode  int32  `json:"errorcode"`
	ErrMsg     string `json:"errmsg"`
}

func (c *Client) AuthCode2Session(ctx context.Context, jscode string) (AuthCode2SessionResponse, error) {
	resp := AuthCode2SessionResponse{}

	path := fmt.Sprintf("/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", c.AppID, c.AppSecret, jscode)

	req, err := c.Sling.New().Get(path).Request()
	if err != nil {
		return resp, err
	}

	_, err = c.Sling.Do(req.WithContext(ctx), &resp, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
