package handler

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/aopontann/qin-todo/common"
)

type TodoList struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Execution_date string `json:"execution_date"`
}

func Pon(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetTodoList(c *gin.Context) {
	db := common.GetDB()
	var (
		id             string
		content        string
		execution_date string
	)
	todo_list := []TodoList{}
	rows, err := db.Query("SELECT id, content, execution_date FROM todo_list")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &content, &execution_date)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, content, execution_date)
		todo_list = append(todo_list, TodoList{Id: id, Content: content, Execution_date: execution_date})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(todo_list)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"items": todo_list,
	})
}
