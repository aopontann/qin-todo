package main

import (
	"database/sql"
)

func GetTodoList(userId string) ([]TodoListInfo, error) {
	var (
		id             string
		content        string
		completed      *sql.NullBool
		execution_date *sql.NullString
	)
	var todoList []TodoListInfo
	rows, err := db.Query("SELECT id, content, completed, execution_date FROM todo_list WHERE user_id = ? AND (completed = 0 OR execution_date IS NULL OR execution_date > now())", userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &content, &completed, &execution_date)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, TodoListInfo{Id: id, Content: content, Completed: completed.Bool, Execution_date: execution_date})
	}
	return todoList, nil
}

func CreateTodo(userId string, todoInfo *PostTodoRequestBody) (string, error) {
	ulid := GetULID()

	if todoInfo.Execution_date == "" {
		_, err := db.Exec("INSERT INTO todo_list (id, content, user_id) VALUES (?,?,?)", ulid, todoInfo.Content, userId)
		if err != nil {
			return "", err
		}
	} else {
		_, err := db.Exec("INSERT INTO todo_list (id, content, execution_date, user_id) VALUES (?,?,?,?)", ulid, todoInfo.Content, todoInfo.Execution_date, userId)
		if err != nil {
			return "", err
		}
	}
	return ulid, nil
}

func UpdateTodo(userId string, todoId string, todoInfo *PutTodoRequestBody) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// リクエストボディにtodoの内容が含まれていた場合、内容の更新処理を行う
	if todoInfo.Content != "" {
		_, err := tx.Exec("UPDATE todo_list SET content = ? WHERE id = ? AND user_id = ? LIMIT 1", todoInfo.Content, todoId, userId)
		if err != nil {
			return err
		}
	}

	// リクエストボディにtodoのやる日が含まれていた場合、やる日の更新処理を行う
	if todoInfo.Execution_date != "" {
		_, err := tx.Exec("UPDATE todo_list SET execution_date = ? WHERE id = ? AND user_id = ? LIMIT 1", todoInfo.Execution_date, todoId, userId)
		if err != nil {
			return err
		}
	}

	// リクエストボディに完了かどうかのbool値が、DBに保存されているcompletedと違う場合、完了状態の更新処理を行う
	reqbComp := 0
	if todoInfo.Completed {
		reqbComp = 1
	}
	_, err = tx.Exec("UPDATE todo_list SET completed = ? WHERE id = ? AND user_id = ? AND NOT completed = ? LIMIT 1", reqbComp, todoId, userId, reqbComp)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
