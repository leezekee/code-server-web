package middleware

import (
	"backend/code"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取cookie
		cookie, err := ctx.Cookie("PHPSESSID")
		if err != nil || cookie == "" {
			response.Fail(ctx, code.NOT_AUTH, "未登录")
			ctx.Abort()
			return
		}
		username := getPHPUserName(cookie)
		if username == nil {
			response.Fail(ctx, code.NOT_AUTH, "未登录")
			ctx.Abort()
			return
		}
		// 用户存在 将user信息写入上下文
		ctx.Set("username", *username)
		ctx.Next()
	}
}

func getPHPUserName(phpSessionId string) *string {
	sessionPah := fmt.Sprintf("/var/lib/php/sessions/sess_%s", phpSessionId)
	if _, err := os.Stat(sessionPah); err != nil {
		return nil
	}
	f, _ := os.Open(sessionPah)
	content := make([]byte, 1024)
	read, _ := f.Read(content)
	contentList := strings.Split(string(content[:read]), ";")
	for _, item := range contentList {
		if strings.HasPrefix(item, "BUCTOJ_user_id") {
			userName := strings.Split(item, "\"")[1]
			return &userName
		}
	}
	return nil
}
