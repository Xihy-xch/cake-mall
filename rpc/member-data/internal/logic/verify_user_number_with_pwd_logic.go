package logic

import (
	"context"

	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	pwdRight = 0
	pwdWrong = 1
)

type VerifyUserNumberWithPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyUserNumberWithPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyUserNumberWithPwdLogic {
	return &VerifyUserNumberWithPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyUserNumberWithPwdLogic) VerifyUserNumberWithPwd(in *member.VerifyUserNumberWithPwdRequest) (*member.VerifyUserNumberWithPwdResponse, error) {
	userPwd, err := l.svcCtx.UserPwdModel.FindOneByNumber(in.GetUserNumber())
	if err != nil && err == sqlx.ErrNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, err
	}

	if userPwd.Password != in.GetPassword() {
		return &member.VerifyUserNumberWithPwdResponse{
			Status: pwdWrong,
		}, nil
	}

	return &member.VerifyUserNumberWithPwdResponse{
		Status: pwdRight,
	}, nil
}
