package main

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var rdb *redis.Client

// Opening a database and save the reference to `Database` struct.
func MysqlInit() {
	// DB接続
	var err error
	db, err = sql.Open("mysql", "user1:pass@tcp(mysql:3306)/qin-todo")
	if err != nil {
		log.Printf("connect error: %s\n", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func RedisInit() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// wget https://dl.google.com/go/go1.18.1.linux-amd64.tar.gz
// sudo tar -xvf go1.18.1.linux-amd64.tar.gz