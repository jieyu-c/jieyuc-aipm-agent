package user_account

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/types"
	user_account_rpc "jieyuc.cn/jieyuc-aipm-agent/rpc/pb/user_account"
)

type GetUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailLogic) GetUserDetail(userId string) (resp *types.GetUserDetailResponse, err error) {
	rpcResp, err := l.svcCtx.UserAccountRpc.GetUserDetail(l.ctx, &user_account_rpc.GetUserDetailReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	if rpcResp.GetUserDetail() == nil {
		return &types.GetUserDetailResponse{}, nil
	}
	ud := rpcResp.GetUserDetail()
	return &types.GetUserDetailResponse{
		UserDetail: types.UserDetail{
			UserId:    ud.GetUserId(),
			Phone:     ud.GetPhone(),
			Username:  ud.GetUsername(),
			Nickname:  ud.GetNickname(),
			Avatar:    ud.GetAvatar(),
			Gender:    ud.GetGender(),
			Status:    ud.GetStatus(),
			CreatedAt: ud.GetCreatedAt(),
			UpdatedAt: ud.GetUpdatedAt(),
		},
	}, nil
}
