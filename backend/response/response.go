package response

import (
	"backend/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, code code.StatusCode, data gin.H, message string) {
	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

func Success(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, code.OK, data, message)
}

func SuccessWithoutData(ctx *gin.Context, message string) {
	Response(ctx, http.StatusOK, code.OK, nil, message)
}

func Fail(ctx *gin.Context, code code.StatusCode, data gin.H, message string) {
	Response(ctx, http.StatusOK, code, data, message)
}

func FailTest(ctx *gin.Context, message string) {
	Response(ctx, http.StatusOK, code.OK, nil, message)
}
