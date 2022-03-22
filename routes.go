package main

import (
	"github.com/aopontann/qin-todo/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", handler.Pon)

	// todoリスト取得機能(デモ版)
	r.GET("/todo_list", handler.GetTodoList)

	auth := r.Group("/auth")
	{
		// google認証画面にリダイレクト
		auth.GET("/google", handler.GoogleAuthRedirect)

		// トークン取得エンドポイント
		auth.GET("/token", handler.GoogleAuthGetToken)
	}

	return r
}
