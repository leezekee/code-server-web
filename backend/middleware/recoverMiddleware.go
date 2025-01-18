package middleware

import (
	"backend/code"
	"backend/response"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, code.SERVER_ERROR, "服务器繁忙，请稍后再试")
			}
		}()

		ctx.Next()
	}
}
