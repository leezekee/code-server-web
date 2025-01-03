package router

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoute(e *gin.Engine) *gin.Engine {
	e.GET("/hello", controller.Hello)

	return e
}
