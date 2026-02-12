package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/pb/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/user-account/internal/svc"
)

type RegisterRpcLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterRpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterRpcLogic {
	return &RegisterRpcLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterRpcLogic) RegisterRpc(in *user_account.RegisterReq) (*user_account.RegisterResp, error) {
	res, err := l.svcCtx.UserAccountApp.Register(l.ctx, in.GetPhone(), in.GetPassword(), in.GetVerifyCode())
	if err != nil {
		l.Logger.Errorf("register rpc failed, err: %v", err)
		return nil, fmt.Errorf("%s", err.Error())
	}
	return &user_account.RegisterResp{
		RegisterStatus: res.OK,
		Msg:            res.Msg,
	}, nil
}
