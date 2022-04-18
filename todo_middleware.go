package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// 指定されたTodoが認証されたユーザーが作成したものかチェックする
func TodoCheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(string)
		todoId := c.Param("todo_id")

		// ユーザーが作成したToDoかチェック
		var id string
		err := db.QueryRow("SELECT id FROM todo_list WHERE id = ? AND user_id = ?", todoId, userId).Scan(&id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(403, gin.H{"error": "Unable to delete a Todo created by another user, or a non-existent Todo ID is specified"})
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
			return
		}

		c.Next()
	}
}
