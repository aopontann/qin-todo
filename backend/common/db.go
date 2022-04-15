package common

import (
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var RDB *redis.Client

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

func RedisInit() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	RDB = rdb
	return RDB
}

func GetRDB() *redis.Client {
	return RDB
}
