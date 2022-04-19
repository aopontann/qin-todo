package main

import (
	"github.com/gin-gonic/gin"
)

func MWGetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid := ""
		// リクエストヘッダーに含まれているセッションIDを取得する
		headers := c.Request.Header
		if len(headers["Session-Id"]) != 0 {
			sid = headers["Session-Id"][0]
		}

		// cookieからセッションIDを取得する
		if sid == "" {
			var err error
			sid, err = c.Cookie("session")
			if err != nil {
				c.Set("userId", "")
				c.Next()
				return
			}
		}

		userid, err := rdb.Get(c, sid).Result()
		if err != nil {
			c.Set("userId", "")
			c.Next()
			return
		}

		c.Set("userId", userid)

		c.Next()
	}
}
