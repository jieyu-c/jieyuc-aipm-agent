package logic

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
	"jieyuc.cn/jieyuc-aipm-agent/service/pb/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/internal/svc"
)

// 统一错误信息，防止通过响应差异枚举用户
const errLoginFailed = "手机号或密码错误"

// 用户状态常量 (与 users 表 status 字段一致)
const (
	userStatusNormal = 0
	userStatusBanned = 1
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
	// 1. 输入校验：手机号必填，密码或验证码至少提供其一
	if in.GetPhone() == "" {
		return nil, fmt.Errorf("手机号不能为空")
	}
	if in.GetPassword() == "" && in.GetVerifyCode() == "" {
		return nil, fmt.Errorf("请输入密码或验证码")
	}

	// 2. 查询用户
	user, err := l.svcCtx.UsersModel.FindOneByPhone(l.ctx, in.GetPhone())
	if err != nil {
		l.Logger.Errorf("login rpc, query db failed, phone: %s, err: %v", in.GetPhone(), err)
		return nil, fmt.Errorf(errLoginFailed)
	}
	if user == nil {
		return nil, fmt.Errorf(errLoginFailed)
	}

	// 3. 检查用户状态（0:正常 1:禁用）
	if user.Status == userStatusBanned {
		return nil, fmt.Errorf("账号已被禁用，请联系客服")
	}

	// 4. 验证凭据：密码登录 或 验证码登录
	if in.GetPassword() != "" {
		// 密码登录：使用 bcrypt 或明文兼容验证
		if !utils.VerifyPassword(user.Password, in.GetPassword()) {
			return nil, fmt.Errorf(errLoginFailed)
		}
	} else {
		// TODO: 验证码登录，需接入验证码服务
		return nil, fmt.Errorf("验证码登录功能暂未开放，请使用密码登录")
	}

	// 5. 生成 JWT
	payload := fmt.Sprintf(`{"userId": "%s", "phone": "%s"}`, user.UserId, user.Phone)
	token, err := utils.GetJwtToken(payload)
	if err != nil {
		l.Logger.Errorf("login rpc, get jwt token failed, err: %v", err)
		return nil, fmt.Errorf("登录失败，请稍后重试")
	}

	return &user_account.LoginResp{
		UserInfo: &user_account.UserInfo{
			UserId: user.UserId,
			Phone:  user.Phone,
		},
		Token: token,
	}, nil
}
