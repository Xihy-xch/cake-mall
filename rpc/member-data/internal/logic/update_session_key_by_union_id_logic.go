package logic

import (
	"context"

	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"
	"cake-mall/rpc/member-data/model"

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

	wechatUser, err := l.svcCtx.WechatUserModel.FindOneByUnionID(in.GetUnionId())

	if err != nil {
		return nil, err
	}
	err = l.svcCtx.WechatUserModel.Update(model.WechatUser{
		Id:         wechatUser.Id,
		Number:     wechatUser.Number,
		UnionId:    wechatUser.UnionId,
		SessionKey: in.GetSessionKey(),
	})
	if err != nil {
		return nil, err
	}
	return &member.UpdateSessionKeyByUnionIdResponse{}, nil
}
