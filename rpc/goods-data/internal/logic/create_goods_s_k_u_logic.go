package logic

import (
	"context"

	"cake-mall/rpc/goods-data/goods"
	"cake-mall/rpc/goods-data/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateGoodsSKULogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsSKULogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsSKULogic {
	return &CreateGoodsSKULogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsSKULogic) CreateGoodsSKU(in *goods.CreateGoodsSKURequest) (*goods.CreateGoodsSKUResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.CreateGoodsSKUResponse{}, nil
}
