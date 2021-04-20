package logic

import (
	"context"
	"time"

	"cake-mall/api/internal/svc"
	"cake-mall/api/internal/types"
	"cake-mall/rpc/member-data/member"

	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	pwdRight     int32 = 0
	pwdWrong     int32 = 1
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginRequest) (*types.LoginResponse, error) {

	resp, err := l.svcCtx.MemberClient.VerifyUserNumberWithPwd(l.ctx, &member.VerifyUserNumberWithPwdRequest{
		UserNumber: req.UserNumber,
		Password:   req.Password,
	})

	if err != nil  && status.Code(err) == codes.NotFound{
		return &types.LoginResponse{
			Code: 1,
			Msg:  "该账户不存在",
		}, nil
	}

	if err != nil {
		return &types.LoginResponse{}, err
	}

	if resp.GetStatus() == pwdWrong {
		return &types.LoginResponse{
			Code: 1,
			Msg:  "用户密码错误请重试",
		}, nil
	}

	return l.generateToken(req.UserNumber)
}

func (l *LoginLogic) generateToken(userNumber int64) (*types.LoginResponse, error) {

	iat := time.Now().Unix()
	secretKey := l.svcCtx.Config.Auth.AccessSecret
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	payloads := make(map[string]interface{}, 1)
	payloads["user_number"] = userNumber

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + accessExpire
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &types.LoginResponse{
			Code: 1,
			Msg:  "生成token失败",
		}, nil
	}
	return &types.LoginResponse{
		Data: types.LoginResponseData{
			AccessToken:  accessToken,
			AccessExpire: iat + accessExpire,
			RefreshAfter: iat + accessExpire/2,
		},
	}, nil
}

