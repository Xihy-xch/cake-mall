package logic

import (
	"context"

	"cake-mall/rpc/goods-data/goods"
	"cake-mall/rpc/goods-data/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateGoodsSPULogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsSPULogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsSPULogic {
	return &CreateGoodsSPULogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsSPULogic) CreateGoodsSPU(in *goods.CreateGoodsSPURequest) (*goods.CreateGoodsSPUResponse, error) {
	// todo: add your logic here and delete this line

	return &goods.CreateGoodsSPUResponse{}, nil
}
