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

	auth := r.Group("/auth")
	{
		// google認証画面にリダイレクト
		auth.GET("/google", handler.GoogleAuthRedirect)

		// トークン取得エンドポイント
		auth.GET("/token", handler.GoogleAuthGetToken)

		auth.POST("/register", handler.UserRegister)

		auth.POST("/login", handler.SessionAuthLogin)

		auth.POST("/logout", handler.SessionAuthLogout)

	}

	user := r.Group("/user")
	{
		user.GET("/", handler.GetUser)
	}

	// 本番環境では使わない検証用パス
	demo := r.Group("/demo")
	{
		// todoリスト取得機能(デモ版)
		demo.GET("/todo_list", handler.GetTodoList)
		demo.GET("/user_hardCode", handler.GetUserHardCode)
		demo.POST("/post_user_demo", handler.PostUserDemo)
		demo.POST("/cookie", handler.CookieDemo)
	}

	return r
}
