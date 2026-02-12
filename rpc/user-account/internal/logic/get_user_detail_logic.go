package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	useraccount "jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/pb/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/rpc/user-account/internal/svc"
)

type GetUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDetailLogic) GetUserDetail(in *user_account.GetUserDetailReq) (*user_account.GetUserDetailResp, error) {
	if in.GetUserId() == "" {
		return nil, fmt.Errorf("userId不能为空")
	}
	user, err := l.svcCtx.UserAccountApp.GetUserDetail(l.ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, useraccount.ErrUserNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		l.Logger.Errorf("GetUserDetail failed, userId: %s, err: %v", in.GetUserId(), err)
		return nil, fmt.Errorf("获取用户详情失败")
	}
	return &user_account.GetUserDetailResp{
		UserDetail: &user_account.UserDetail{
			UserId:    user.UserId,
			Phone:     user.Phone,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Gender:    user.Gender,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
		},
	}, nil
}
