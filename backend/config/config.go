package config

import (
	"github.com/spf13/viper"
	"os"
)

// 初始化配置文件
func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir)
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
