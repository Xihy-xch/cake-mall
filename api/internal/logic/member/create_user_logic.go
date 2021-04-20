package logic

import (
	"context"

	"cake-mall/api/internal/svc"
	"cake-mall/api/internal/types"
	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateUserLogic {
	return CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req types.CreateUserRequest) (*types.CreateUserResponse, error) {
	_, err := l.svcCtx.MemberClient.CreateUser(l.ctx, &member.CreateUserRequest{
		UserNumber: req.UserNumber,
		Password:   req.Password,
		Mobile:     req.Mobile,
		UserName:   req.UserName,
	})
	if err != nil {
		return &types.CreateUserResponse{
			Code: 1,
			Msg:  "创建用户失败",
		}, err
	}

	return &types.CreateUserResponse{}, nil
}
