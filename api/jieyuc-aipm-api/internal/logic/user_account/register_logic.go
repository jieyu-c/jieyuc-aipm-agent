// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user_account

import (
	"context"
	"strings"

	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/types"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/pb/user_account"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 参数预处理：去除首尾空格
	req.Phone = strings.TrimSpace(req.Phone)
	req.VerifyCode = strings.TrimSpace(req.VerifyCode)
	rpcResp, err := l.svcCtx.UserAccountRpc.RegisterRpc(l.ctx, &user_account.RegisterReq{
		Phone:      req.Phone,
		Password:   req.Password,
		VerifyCode: req.VerifyCode,
	})
	if err != nil {
		l.Logger.Errorf("register: rpc call failed, err: %v", err)
		return &types.RegisterResponse{
			RegisterStatus: false,
			Msg:            "系统繁忙，请稍后重试",
		}, nil
	}

	return &types.RegisterResponse{
		RegisterStatus: rpcResp.RegisterStatus,
		Msg:            rpcResp.Msg,
	}, nil
}
