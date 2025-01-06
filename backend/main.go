package main

import (
	_ "backend/config" // 引入配置文件执行init函数
	_ "backend/db"     // 引入数据库执行init函数
	"backend/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	server := gin.Default()
	server = router.CollectRoute(server)
	port := viper.GetString("server.port")
	if port != "" {
		panic(server.Run(":" + port))
	}
	panic(server.Run())
}
