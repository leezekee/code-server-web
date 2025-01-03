package main

import (
	"backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: 初始化配置、连接数据库

	server := gin.Default()
	server = router.CollectRoute(server)
	panic(server.Run(":8080"))
}
