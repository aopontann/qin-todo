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
	
	// middlewareで認証をして成功すると、ここでユーザーIDを取得できる
	userId := c.MustGet("userId").(string)

	// MySQLに保存されているユーザー情報を取得する
	db := common.GetDB()
	err := db.QueryRow("SELECT id, name, email, avatar_url FROM users WHERE id = ?", userId).Scan(&id, &name, &email, &avatar_url)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// この記述よくなさそうだから、他にいい方法があるか調べてみる
	if avatar_url != nil {
		c.JSON(200, gin.H{
			"id":    id,
			"name":  name,
			"email": email,
			"avatar_url": avatar_url.String,
		})
		return
	}
	c.JSON(200, gin.H{
		"id":    id,
		"name":  name,
		"email": email,
		"avatar_url": nil,
	})

}
