package logic

import (
	"context"
	"fmt"
	"strings"

	"cake-mall/rpc/member-data/internal/model"
	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
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
	conn := sqlx.NewMysql(l.svcCtx.Config.DB.Member)

	err := conn.Transact(func(session sqlx.Session) error {

		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", "`user`",
			strings.Join(stringx.Remove(builderx.RawFieldNames(&model.User{}), "`id`", "`create_time`", "`update_time`"), ","))
		_, err := session.Exec(query, in.GetUserNumber(), in.GetUserName(), in.Mobile)
		if err != nil {
			return err
		}

		query = fmt.Sprintf("insert into %s (%s) values (?, ?)", "`user_pwd`",
			strings.Join(stringx.Remove(builderx.RawFieldNames(&model.UserPwd{}), "`id`", "`create_time`", "`update_time`"), ","))
		_, err = session.Exec(query, in.GetUserNumber(), in.GetPassword())
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &member.CreateUserResponse{}, nil
}
