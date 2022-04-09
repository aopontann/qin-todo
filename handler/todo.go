package handler

import (
	"database/sql"

	"github.com/aopontann/qin-todo/common"
	"github.com/gin-gonic/gin"
)

type TodoListInfo struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Completed      bool `json:"completed"`
	Execution_date *sql.NullString `json:"execution_date"`
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
