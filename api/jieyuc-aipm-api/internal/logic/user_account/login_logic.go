// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user_account

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/types"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/pb/user_account"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1. 请求参数校验
	if req.Phone == "" {
		return nil, fmt.Errorf("手机号不能为空")
	}
	if req.Password == "" && req.VerifyCode == "" {
		return nil, fmt.Errorf("请输入密码或验证码")
	}

	// 2. 调用 RPC 登录
	rpcResp, err := l.svcCtx.UserAccountRpc.LoginRpc(l.ctx, &user_account.LoginReq{
		Phone:      req.Phone,
		Password:   req.Password,
		VerifyCode: req.VerifyCode,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		UserInfo: types.UserInfo{
			Phone:  rpcResp.UserInfo.Phone,
			UserId: rpcResp.UserInfo.UserId,
		},
		Token: rpcResp.Token,
	}, nil
}
