package logic

import (
	"context"

	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateSessionKeyByUnionIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSessionKeyByUnionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSessionKeyByUnionIdLogic {
	return &UpdateSessionKeyByUnionIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSessionKeyByUnionIdLogic) UpdateSessionKeyByUnionId(in *member.UpdateSessionKeyByUnionIdRequest) (*member.UpdateSessionKeyByUnionIdResponse, error) {
	// todo: add your logic here and delete this line

	return &member.UpdateSessionKeyByUnionIdResponse{}, nil
}
