package main

import (
	"github.com/gin-gonic/gin"
)

func GetTodoHandler(c *gin.Context) {
	// middlewareで認証をして成功すると、ここでユーザーIDを取得できる
	userId := c.MustGet("userId").(string)

	todoList, err := GetTodoList(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"items": todoList,
	})
}

func PostTodoHandler(c *gin.Context) {
	var reqb PostTodoRequestBody
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId").(string)
	var ulid string
	ulid, err = CreateTodo(userId, &reqb)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"id": ulid, "content": reqb.Content, "execution_date": nil, "user_id": userId})
}

func PutTodoHandler(c *gin.Context) {
	var reqb PutTodoRequestBody
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId").(string)
	todoId := c.Param("todo_id")

	err = UpdateTodo(userId, todoId, &reqb)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}

func DeleteTodoHandler(c *gin.Context) {
	userId := c.MustGet("userId").(string)
	todoId := c.Param("todo_id")

	_, err := db.Exec("DELETE FROM todo_list WHERE id = ? AND user_id = ? LIMIT 1", todoId, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}
