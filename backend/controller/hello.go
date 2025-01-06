package controller

import (
	"backend/db"
	"backend/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {

	rdb := db.GetRedis()
	//fmt.Println(rdb)
	//rdb.Set(ctx, "key", "value", 0)
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		//response.FailTest(ctx, "redis get failed" + err)
		fmt.Println(err)
		response.FailTest(ctx, "redis get failed")
		return
	}

	//response.SuccessWithoutData(ctx, "put success")
	response.Success(ctx, gin.H{"val": val}, "redis get success")
}
