package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"jieyuc.cn/jieyuc-aipm-agent/internal/model/users"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
	"jieyuc.cn/jieyuc-aipm-agent/service/pb/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/service/user-account/internal/svc"
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
	// 2. 验证码校验
	// TODO: 后续接入验证码服务，当前使用临时验证码用于开发测试
	if !l.verifyCode(in.GetPhone(), in.GetVerifyCode()) {
		return &user_account.RegisterResp{
			RegisterStatus: false,
			Msg:            "验证码错误或已过期",
		}, nil
	}

	// 3. 检查手机号是否已注册
	exists, err := l.svcCtx.UsersModel.Contains(l.ctx, in.GetPhone())
	if err != nil {
		l.Logger.Errorf("register: check phone exists failed, phone: %s, err: %v",
			maskPhone(in.GetPhone()), err)
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}

	if exists {
		// 出于安全考虑，不明确告知手机号已注册
		return &user_account.RegisterResp{
			RegisterStatus: false,
			Msg:            "注册失败，请检查输入信息或联系客服",
		}, nil
	}

	// 4. 密码加密
	hashedPassword, err := utils.HashPassword(in.GetPassword())
	if err != nil {
		l.Logger.Errorf("register: hash password failed, err: %v", err)
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}

	// 5. 生成用户ID并入库
	userId := generateUserId()
	_, err = l.svcCtx.UsersModel.Insert(l.ctx, &users.Users{
		UserId:   userId,
		Phone:    in.GetPhone(),
		Username: userId, // 默认用户名使用 userId
		Password: hashedPassword,
	})

	if err != nil {
		l.Logger.Errorf("register: insert user failed, phone: %s, err: %v",
			maskPhone(in.GetPhone()), err)
		return nil, fmt.Errorf("系统繁忙，请稍后重试")
	}

	l.Logger.Infof("register: success, userId: %s, phone: %s", userId, maskPhone(in.GetPhone()))

	return &user_account.RegisterResp{
		RegisterStatus: true,
		Msg:            "注册成功",
	}, nil
}

// verifyCode 校验验证码
// TODO: 接入 Redis 或验证码服务进行真实校验
func (l *RegisterRpcLogic) verifyCode(phone, code string) bool {
	// 开发测试阶段，使用临时验证码
	// 正式环境应从 Redis 或验证码服务获取验证
	const devTestCode = "123456"
	return code == devTestCode
}

// generateUserId 生成用户唯一ID
// 使用 UUID v4，保证全局唯一且不可预测
func generateUserId() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// maskPhone 手机号脱敏，用于日志输出
// 例如：13812345678 -> 138****5678
func maskPhone(phone string) string {
	if len(phone) != 11 {
		return "***"
	}
	return phone[:3] + "****" + phone[7:]
}
