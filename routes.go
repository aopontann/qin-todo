package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	// すべてのオリジンを許可する(本番環境にデプロイするまでにちゃんと設定する)
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/ping", PonHandler)

	auth := r.Group("/auth")
	{
		// google認証画面にリダイレクト
		auth.GET("/google", GoogleAuthRedirectHandler)

		// トークン取得エンドポイント
		auth.GET("/token", GoogleAuthGetTokenHandler)

		auth.POST("/register", UserRegisterHandler)

		auth.POST("/login", SessionAuthLoginHandler)

		auth.POST("/logout", SessionAuthLogoutHandler)

	}

	user := r.Group("/users")
	{
		user.Use(MWGetUserID())
		user.GET("/", GetUserHandler)
		user.PUT("/", PutUserHandler)
	}

	todo := r.Group("/todos")
	{
		todo.Use(MWGetUserID())
		todo.GET("/", GetTodoHandler)
		todo.POST("/", PostTodoHandler)
		todoPutDelete := todo.Group("/")
		{
			todoPutDelete.Use(TodoCheckUser())
			todoPutDelete.PUT("/:todo_id", PutTodoHandler)
			todoPutDelete.DELETE("/:todo_id", DeleteTodoHandler)
		}
	}

	// 本番環境では使わない検証用パス
	demo := r.Group("/demo")
	{
		// todoリスト取得機能(デモ版)
		demo.GET("/todo_list", GetTodoListHandler)
		demo.GET("/user_hardCode", GetUserHardCodeHandler)
		demo.POST("/post_user_demo", PostUserDemoHandler)
		demo.POST("/cookie", CookieDemoHandler)
	}

	return r
}
