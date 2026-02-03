package user_account

import (
	"encoding/json"
	"net/http"
	"strings"

	restful "github.com/jieyu-c/jieyuc-common/types"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/logic/user_account"
	"jieyuc.cn/jieyuc-aipm-agent/api/jieyuc-aipm-api/internal/svc"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
)

// writeUnauthorized 按 REST 规范返回 401，与 restful 用法一致（JSON 体 + 状态码）
func writeUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func GetUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(r.Header.Get("Authorization"))
		if token == "" {
			writeUnauthorized(w, "请先登录")
			return
		}
		payload, err := utils.ParseJwtToken(token)
		if err != nil {
			writeUnauthorized(w, "token 无效或已过期")
			return
		}
		var claims struct {
			UserId string `json:"userId"`
		}
		if err := json.Unmarshal([]byte(payload), &claims); err != nil || claims.UserId == "" {
			writeUnauthorized(w, "token 无效")
			return
		}

		l := user_account.NewGetUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetUserDetail(claims.UserId)
		restful.Response(w, resp, err)
	}
}
