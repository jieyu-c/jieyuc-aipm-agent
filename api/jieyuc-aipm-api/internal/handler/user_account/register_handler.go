// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user_account

import (
	"net/http"

	restful "github.com/jieyu-c/jieyuc-common/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/logic/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user_account.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		restful.Response(w, resp, err)

	}
}
