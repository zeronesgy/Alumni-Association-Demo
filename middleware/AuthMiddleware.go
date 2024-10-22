package middleware

import (
	"Alumni-Association-Demo/common"
	"Alumni-Association-Demo/model"
	"Alumni-Association-Demo/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware 用户中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token format 若为空，或不以"Bearer "开头，说明token错误
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		// 验证通过获取 claims 中的userId
		userId := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		// 用户
		if user.ID == 0 {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "权限不足")
			ctx.Abort()
			return
		}

		// 用户存在 将user的信息写入上下文
		ctx.Set("user", user)

		ctx.Next()
	}
}
