package handler

import (
	"database/sql"

	"github.com/aopontann/qin-todo/common"
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

func GetTodo(c *gin.Context) {
	var (
		id             string
		content        string
		completed      *sql.NullBool
		execution_date *sql.NullString
	)
	var todoList []TodoListInfo

	// middlewareで認証をして成功すると、ここでユーザーIDを取得できる
	userId := c.MustGet("userId").(string)

	// MySQLに保存されているToDo情報を取得する
	db := common.GetDB()
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

func PostTodo(c *gin.Context) {
	var reqb TodoRequestBody
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId").(string)
	ulid := common.GetULID()
	db := common.GetDB()

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

func PutTodo(c *gin.Context) {
	var reqb PutTodoRequestBody
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId").(string)
	todoId := c.Param("todo_id")

	tx, err := common.DB.Begin()
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

	var completed int
	err = tx.QueryRow("SELECT completed FROM todo_list WHERE id = ? AND user_id = ?", todoId, userId).Scan(&completed)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	reqbComp := 0
	if reqb.Completed {
		reqbComp = 1
	}
	if completed != reqbComp {
		_, err := tx.Exec("UPDATE todo_list SET completed = ? WHERE id = ? AND user_id = ? LIMIT 1", reqbComp, todoId, userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Status(200)

}
