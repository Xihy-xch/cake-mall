// Code generated by goctl. DO NOT EDIT!
// Source: member_data.proto

//go:generate mockgen -destination ./member_mock.go -package memberclient -source $GOFILE

package memberclient

import (
	"context"

	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	VerifyUserNumberWithPwdRequest    = member.VerifyUserNumberWithPwdRequest
	VerifyUserNumberWithPwdResponse   = member.VerifyUserNumberWithPwdResponse
	UpdateSessionKeyByUnionIdRequest  = member.UpdateSessionKeyByUnionIdRequest
	UpdateSessionKeyByUnionIdResponse = member.UpdateSessionKeyByUnionIdResponse

	Member interface {
		UpdateSessionKeyByUnionId(ctx context.Context, in *UpdateSessionKeyByUnionIdRequest) (*UpdateSessionKeyByUnionIdResponse, error)
		VerifyUserNumberWithPwd(ctx context.Context, in *VerifyUserNumberWithPwdRequest) (*VerifyUserNumberWithPwdResponse, error)
	}

	defaultMember struct {
		cli zrpc.Client
	}
)

func NewMember(cli zrpc.Client) Member {
	return &defaultMember{
		cli: cli,
	}
}

func (m *defaultMember) UpdateSessionKeyByUnionId(ctx context.Context, in *UpdateSessionKeyByUnionIdRequest) (*UpdateSessionKeyByUnionIdResponse, error) {
	client := member.NewMemberClient(m.cli.Conn())
	return client.UpdateSessionKeyByUnionId(ctx, in)
}

func (m *defaultMember) VerifyUserNumberWithPwd(ctx context.Context, in *VerifyUserNumberWithPwdRequest) (*VerifyUserNumberWithPwdResponse, error) {
	client := member.NewMemberClient(m.cli.Conn())
	return client.VerifyUserNumberWithPwd(ctx, in)
}
