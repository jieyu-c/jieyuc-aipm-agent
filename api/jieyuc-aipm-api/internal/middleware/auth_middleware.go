// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"jieyuc.cn/jieyuc-aipm-agent/internal/domain/useraccount"
	"jieyuc.cn/jieyuc-aipm-agent/internal/utils"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// writeUnauthorized 按 REST 规范返回 401，与 restful 用法一致（JSON 体 + 状态码）
func writeUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header
		// Get auth token from headers
		authToken := headers.Get("Authorization")
		// Parse jwt token to payload string
		payload, err := utils.ParseJwtToken(authToken)
		if err != nil {
			// if err, response 401
			writeUnauthorized(w, "Auth token verify failed")
			return
		}
		// parse to TokenPayload object
		tokenPayload, err := useraccount.Json2TokenPayload(payload)
		httpCtx := r.Context()
		ctx := context.WithValue(httpCtx, utils.ContextKeyUserId, tokenPayload.UserId)
		next(w, r.WithContext(ctx))
	}
}
