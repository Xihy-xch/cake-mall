package logic

import (
	"context"

	"cake-mall/api/internal/svc"
	"cake-mall/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateWechatSessionKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateWechatSessionKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateWechatSessionKeyLogic {
	return UpdateWechatSessionKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	authCode2SessionSuccess = 0
)

func (l *UpdateWechatSessionKeyLogic) UpdateWechatSessionKey(req types.UpdateWechatSessionKeyRequest) (*types.UpdateWechatSessionKeyResponse, error) {
	authCode2SessionResponse, err := l.svcCtx.WechatClient.AuthCode2Session(l.ctx, req.JsCode)
	if err != nil {
		return &types.UpdateWechatSessionKeyResponse{
			Code: 1,
			Msg:  "微信接口调用失败",
		}, err
	}
	
	if authCode2SessionResponse.ErrorCode != authCode2SessionSuccess {
		return &types.UpdateWechatSessionKeyResponse{
			Code: 1,
			Msg:  authCode2SessionResponse.ErrMsg,
		}, nil
	}



	return &types.UpdateWechatSessionKeyResponse{}, nil
}
