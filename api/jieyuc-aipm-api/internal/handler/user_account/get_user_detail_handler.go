package user_account

import (
	"net/http"

	restful "github.com/jieyu-c/jieyuc-common/types"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/logic/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
)

func GetUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user_account.NewGetUserDetailLogic(r.Context(), svcCtx)
		ctx := r.Context()
		userId := utils.ContextKeyUserId.GetStringValueFromContext(ctx)
		resp, err := l.GetUserDetail(userId)
		restful.Response(w, resp, err)
	}
}
