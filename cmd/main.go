package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// .envの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	r := InitRouter()

	r.Run()
}
