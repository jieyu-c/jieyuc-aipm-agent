package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/pb/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/user-account/internal/svc"
)

type LoginRpcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginRpcLogic {
	return &LoginRpcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginRpcLogic) LoginRpc(in *user_account.LoginReq) (*user_account.LoginResp, error) {
	res, err := l.svcCtx.UserAccountApp.Login(l.ctx, in.GetPhone(), in.GetPassword(), in.GetVerifyCode())
	if err != nil {
		l.Logger.Errorf("login rpc failed, phone: %s, err: %v", in.GetPhone(), err)
		return nil, fmt.Errorf("%s", err.Error())
	}
	return &user_account.LoginResp{
		UserInfo: &user_account.UserInfo{
			UserId: res.User.UserId,
			Phone:  res.User.Phone,
		},
		Token: res.Token,
	}, nil
}
