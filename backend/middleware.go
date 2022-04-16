package main

import (
	"github.com/aopontann/qin-todo/backend/common"
	"github.com/gin-gonic/gin"
)

func MWGetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cookieからセッションIDを取得する
		sid, err := c.Cookie("session")
		if err != nil {
			c.JSON(401, gin.H{"error": "session invailed"})
			return
		}

		// redisに保存されているユーザーIDを取得する
		rdb := common.GetRDB()
		userid, err := rdb.Get(c, sid).Result()
		if err != nil {
			c.JSON(401, gin.H{"error": "session invailed!"})
			return
		}

		c.Set("userId", userid)

		c.Next()
	}
}
