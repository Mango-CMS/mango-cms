package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Mango-CMS/mango-cms/internal/auth"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 从请求头中获取token
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				next.ServeHTTP(w, r)
				return
			}

			// 解析Bearer token
			parts := strings.Split(authorization, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				next.ServeHTTP(w, r)
				return
			}

			// 验证token
			claims, err := auth.ParseToken(parts[1])
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// 将用户信息添加到上下文中
			ctx := context.WithValue(r.Context(), "user", claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// RequireAuth 需要认证的指令
func RequireAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, ok := ctx.Value("user").(*auth.Claims)
	if !ok {
		return nil, &gqlerror.Error{
			Message: "请先登录",
			Path:    graphql.GetPath(ctx),
		}
	}

	return next(ctx)
}

// RequireRole 需要特定角色的指令
func RequireRole(role string) func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		claims, ok := ctx.Value("user").(*auth.Claims)
		if !ok {
			return nil, &gqlerror.Error{
				Message: "请先登录",
				Path:    graphql.GetPath(ctx),
			}
		}

		if claims.Role != role && claims.Role != "admin" {
			return nil, &gqlerror.Error{
				Message: "权限不足",
				Path:    graphql.GetPath(ctx),
			}
		}

		return next(ctx)
	}
}
