package main

import (
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

		userid, err := rdb.Get(c, sid).Result()
		if err != nil {
			c.JSON(401, gin.H{"error": "session invailed!"})
			return
		}

		c.Set("userId", userid)

		c.Next()
	}
}
