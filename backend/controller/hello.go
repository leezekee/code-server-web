package controller

import (
	"backend/response"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	response.SuccessWithoutData(ctx, "Hello, World!")
}
