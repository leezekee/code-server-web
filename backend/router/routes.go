package router

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(e *gin.Engine) *gin.Engine {
	e.GET("/hello", controller.Hello)

	userRoute := e.Group("/user")
	userRoute.Use(middleware.AuthMiddleware())
	userRoute.POST("/remove", controller.RemoveUser)
	return e
}
