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

func TestVerifyUserNumberWithPwdLogic_VerifyUserNumberWithPwd(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *member.VerifyUserNumberWithPwdRequest
	}
	var configFile = flag.String("f", "/home/xch/go/src/cake-mall/rpc/member-data/etc/member_data.yaml", "the config file")
	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *member.VerifyUserNumberWithPwdResponse
		wantErr bool
	}{
		{
			name: "查询不到账户",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svcCtx,
			},
			args: args{in: &member.VerifyUserNumberWithPwdRequest{
				UserNumber: 12345678,
				Password:   "pwd123",
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &VerifyUserNumberWithPwdLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.VerifyUserNumberWithPwd(tt.args.in)
			t.Log(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyUserNumberWithPwd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyUserNumberWithPwd() got = %v, want %v", got, tt.want)
			}
		})
	}
}
