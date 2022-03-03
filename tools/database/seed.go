package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user1:pass@tcp(mysql:3306)/qin-todo")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Openエラー")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Pingエラー")
	}
	// デモユーザーデータを追加
	stmt, err := db.Prepare("INSERT INTO users(id, name, email, password) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("2", "太郎", "def@example.com", "123")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d\n", lastId)
	// デモtodoデータを追加
	stmtTodo, err := db.Prepare("INSERT INTO todo_list(id, content, execution_date, user_id) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	resTodo, err := stmtTodo.Exec("2", "テスト", "2019-10-04 15:25:07", "1")
	if err != nil {
		log.Fatal(err)
	}
	todoLastId, err := resTodo.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("todoID = %d\n", todoLastId)
}
