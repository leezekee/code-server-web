package utils

import (
	"backend/db"
	"github.com/gin-gonic/gin"
)

const (
	User2Port     = "UserPort"
	User2Token    = "UserToekn"
	Port2User     = "PortUser"
	PortStartTime = "PortStartTime"
	FileSaveMode  = "FileSaveMode"
)

func DeleteUserFromRedis(ctx *gin.Context, username string) error {
	rdc := db.GetRedis()
	_, ok := rdc.Get(ctx, User2Port+username).Result()
	if ok != nil {
		return ok
	}
	port, _ := rdc.Get(ctx, User2Port+username).Result()
	token, ok := rdc.Get(ctx, User2Token+username).Result()

	saveUserFile()
	releasePod()

	if ok == nil {
		rdc.Del(ctx, token)
	}
	rdc.Del(ctx, User2Port+username)
	rdc.Del(ctx, User2Token+username)
	rdc.Del(ctx, Port2User+port)
	rdc.Del(ctx, PortStartTime+port)
	rdc.Del(ctx, FileSaveMode+username)

	return nil
}
