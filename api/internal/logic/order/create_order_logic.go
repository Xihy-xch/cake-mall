package logic

import (
	"context"

	"cake-mall/api/internal/svc"
	"cake-mall/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateOrderLogic {
	return CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req types.CreateOrderRequest) (*types.CreateOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &types.CreateOrderResponse{
		Code: 0,
		Msg: "success",
	}, nil
}
