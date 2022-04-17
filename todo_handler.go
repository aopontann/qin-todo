package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type TodoListInfo struct {
	Id             string          `json:"id"`
	Content        string          `json:"content"`
	Completed      bool            `json:"completed"`
	Execution_date *sql.NullString `json:"execution_date"`
}

type TodoRequestBody struct {
	Content        string `json:"content"`
	Execution_date string `json:"execution_date"`
}

type PutTodoRequestBody struct {
	Content        string `json:"content"`
	Completed      bool   `json:"completed,omitempty"`
	Execution_date string `json:"execution_date"`
}

func GetTodoHandler(c *gin.Context) {
	var (
		id             string
		content        string
		completed      *sql.NullBool
		execution_date *sql.NullString
	)
	var todoList []TodoListInfo

	// middlewareで認証をして成功すると、ここでユーザーIDを取得できる
	userId := c.MustGet("userId").(string)

	rows, err := db.Query("SELECT id, content, completed, execution_date FROM todo_list WHERE user_id = ? AND (completed = 0 OR execution_date IS NULL OR execution_date > now())", userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &content, &completed, &execution_date)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		todoList = append(todoList, TodoListInfo{Id: id, Content: content, Completed: completed.Bool, Execution_date: execution_date})
	}

	c.JSON(200, gin.H{
		"items": todoList,
	})
}

func PostTodoHandler(c *gin.Context) {
	var reqb TodoRequestBody
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId").(string)
	ulid := GetULID()

	if reqb.Execution_date == "" {
		_, err := db.Exec("INSERT INTO todo_list (id, content, user_id) VALUES (?,?,?)", ulid, reqb.Content, userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"id": ulid, "content": reqb.Content, "execution_date": nil, "user_id": userId})

	} else {
		_, err := db.Exec("INSERT INTO todo_list (id, content, execution_date, user_id) VALUES (?,?,?,?)", ulid, reqb.Content, reqb.Execution_date, userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.Status(201)
	}
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

	// ユーザーが作成したToDoかチェック
	var id string
	err = db.QueryRow("SELECT id FROM todo_list WHERE id = ? AND user_id = ?", todoId, userId).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(403, gin.H{"error": "Todos created by other users cannot be updated, or a non-existent todo ID is specified"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer tx.Rollback()

	if reqb.Content != "" {
		_, err := tx.Exec("UPDATE todo_list SET content = ? WHERE id = ? AND user_id = ? LIMIT 1", reqb.Content, todoId, userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	if reqb.Execution_date != "" {
		_, err := tx.Exec("UPDATE todo_list SET execution_date = ? WHERE id = ? AND user_id = ? LIMIT 1", reqb.Execution_date, todoId, userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	reqbComp := 0
	if reqb.Completed {
		reqbComp = 1
	}
	_, err = tx.Exec("UPDATE todo_list SET completed = ? WHERE id = ? AND user_id = ? AND NOT completed = ? LIMIT 1", reqbComp, todoId, userId, reqbComp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)

}

func DeleteTodoHandler(c *gin.Context) {
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

	_, err = db.Exec("DELETE FROM todo_list WHERE id = ? AND user_id = ? LIMIT 1", todoId, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)
}
