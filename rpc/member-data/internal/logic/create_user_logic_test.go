package logic

import (
	"context"
	"flag"
	"reflect"
	"testing"

	"cake-mall/rpc/member-data/internal/config"
	"cake-mall/rpc/member-data/internal/svc"
	"cake-mall/rpc/member-data/member"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
)

func TestCreateUserLogic_CreateUser(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *member.CreateUserRequest
	}

	var configFile = flag.String("f", "/home/xch/go/src/cake-mall/rpc/member-data/etc/member_data.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *member.CreateUserResponse
		wantErr bool
	}{
		{
			name:    "创建用户",
			fields:  fields{
				ctx:    context.Background(),
				svcCtx: svcCtx,
			},
			args:    args{
				in: &member.CreateUserRequest{
					UserNumber: 0,
					Password:   "pwd1234",
					Mobile:     "15694513786",
					UserName:   "xiaoMing",
				},
			},
			want:    &member.CreateUserResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &CreateUserLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.CreateUser(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
