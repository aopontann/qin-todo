package main

import "database/sql"

type TodoListInfo struct {
	Id             string          `json:"id"`
	Content        string          `json:"content"`
	Completed      bool            `json:"completed"`
	Execution_date *sql.NullString `json:"execution_date"`
}

type PostTodoRequestBody struct {
	Content        string `json:"content"`
	Execution_date string `json:"execution_date"`
}

type PutTodoRequestBody struct {
	Content        string `json:"content"`
	Completed      bool   `json:"completed,omitempty"`
	Execution_date string `json:"execution_date"`
}
