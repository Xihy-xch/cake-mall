package logic

import (
	"context"

	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"
	"cake-mall/rpc/member-data/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *member.CreateUserRequest) (*member.CreateUserResponse, error) {
	if in.GetUserNumber() == 0 {
		in.UserNumber = l.svcCtx.IDGenerator.GetNumber()
	}
	//Todo 事物操作数据
	_, err := l.svcCtx.UserModel.Insert(model.User{
		Number:     in.GetUserNumber(),
		UserName:   in.GetUserName(),
		Mobile:     in.GetMobile(),
	})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.UserPwdModel.Insert(model.UserPwd{
		Number:   in.GetUserNumber(),
		Password: in.GetPassword(),
	})

	return &member.CreateUserResponse{}, nil
}
