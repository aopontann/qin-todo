package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func mysqlDemo() {
	db, err := sql.Open("mysql","user1:pass@tcp(mysql:3306)/qin-todo")
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
}
