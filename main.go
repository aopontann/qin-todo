package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/aopontann/qin-todo/common"
)

func main() {
	// .envの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Google認証に必要な設定
	common.GoogleAuthInit()

	// DBの初期化
	db := common.Init()
	// プールの切断
	defer db.Close()

	r := InitRouter()

	r.Run()
}
