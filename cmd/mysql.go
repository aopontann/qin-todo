package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func mysqlDemo() {
	db, err := sql.Open("mysql", "user1:pass@tcp(mysql:3306)/qin-todo")
	// db, err := sql.Open("mysql","user1:pass@tcp(127.0.0.1:3306)/qin-todo") //ホストPC上から接続する場合
	if err != nil {
		log.Fatal(err)
		fmt.Println("Openエラー")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Pingエラー")
	}

	var (
		id string
		content string
		execution_date string
	)

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
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
