package db

import (
	//"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	//"github.com/spf13/viper"
)

var _redis *redis.Client

func init() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	password := viper.GetString("redis.password")
	addr := fmt.Sprintf("%s:%s", host, port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // 没有密码，默认值
		DB:       0,        // 默认DB 0
	})
	_redis = rdb
}

func GetRedis() *redis.Client {
	return _redis
}
