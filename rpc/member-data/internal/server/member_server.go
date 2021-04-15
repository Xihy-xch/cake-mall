// Code generated by goctl. DO NOT EDIT!
// Source: member_data.proto

package server

import (
	"context"

	"cake-mall/rpc/member-data/internal/logic"
	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"
)

type MemberServer struct {
	svcCtx *svc.ServiceContext
}

func NewMemberServer(svcCtx *svc.ServiceContext) *MemberServer {
	return &MemberServer{
		svcCtx: svcCtx,
	}
}

func (s *MemberServer) UpdateSessionKeyByUnionId(ctx context.Context, in *member.UpdateSessionKeyByUnionIdRequest) (*member.UpdateSessionKeyByUnionIdResponse, error) {
	l := logic.NewUpdateSessionKeyByUnionIdLogic(ctx, s.svcCtx)
	return l.UpdateSessionKeyByUnionId(in)
}

func (s *MemberServer) VerifyUserNumberWithPwd(ctx context.Context, in *member.VerifyUserNumberWithPwdRequest) (*member.VerifyUserNumberWithPwdResponse, error) {
	l := logic.NewVerifyUserNumberWithPwdLogic(ctx, s.svcCtx)
	return l.VerifyUserNumberWithPwd(in)
}
