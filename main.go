package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// .envの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Google認証に必要な設定
	GoogleAuthInit()

	// DBの初期化
	MysqlInit()
	// プールの切断
	defer db.Close()

	RedisInit()

	r := InitRouter()

	r.Run()
}
