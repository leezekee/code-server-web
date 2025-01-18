package controller

import (
	"backend/code"
	"backend/response"
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func RemoveUser(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	err := utils.DeleteUserFromRedis(ctx, *username.(*string))
	if err != nil {
		response.Fail(ctx, code.RELEASE_ERROR, "服务器繁忙，请稍后再试")
	}
	response.SuccessRemoveUser(ctx)
}
