package common

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Opening a database and save the reference to `Database` struct.
func Init() *sql.DB {
	// DB接続
	db, err := sql.Open("mysql", "user1:pass@tcp(mysql:3306)/qin-todo")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	return DB
}

func GetDB() *sql.DB {
	return DB
}
