package handler

import (
	"database/sql"

	"github.com/aopontann/qin-todo/common"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var (
		id         string
		name       string
		email      string
		avatar_url *sql.NullString
	)
	// cookieからセッションIDを取得する
	sid, err := c.Cookie("session")
	if err != nil {
		c.JSON(401, gin.H{"error": "session invailed"})
		return
	}

	// redisに保存されているユーザーIDを取得する
	rdb := common.GetRDB()
	usrid, err := rdb.Get(c, sid).Result()
	if err != nil {
		c.JSON(401, gin.H{"error": "session invailed!"})
		return
	}

	// MySQLに保存されているユーザー情報を取得する
	db := common.GetDB()
	err = db.QueryRow("SELECT id, name, email, avatar_url FROM users WHERE id = ?", usrid).Scan(&id, &name, &email, &avatar_url)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":    id,
		"name":  name,
		"email": email,
		"avatar_url": avatar_url,
	})
}
