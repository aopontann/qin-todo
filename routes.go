package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Session-Id"},
		ExposeHeaders: []string{"Content-Type", "Content-Length"},
		AllowCredentials: true,
		// MaxAge: 24 * time.Hour,
	}))

	r.GET("/ping", PonHandler)

	auth := r.Group("/auth")
	{
		// google認証画面にリダイレクト
		auth.GET("/google", GoogleAuthRedirectHandler)

		// トークン取得エンドポイント
		auth.GET("/token", GoogleAuthGetTokenHandler)

		// ユーザー登録エンドポイント
		auth.POST("/register", UserRegisterHandler)

		// メールアドレスとパスワードでログイン
		auth.POST("/login", SessionAuthLoginHandler)

		// ログアウト
		auth.POST("/logout", SessionAuthLogoutHandler)
	}

	user := r.Group("/users")
	{
		// sessionIDを使ってredisからユーザーIDを取得する
		user.Use(MWGetUserID())

		// ユーザー情報を取得する
		user.GET("", GetUserHandler)

		// ユーザーの名前とアイコンを変更する(実装中)
		user.PUT("", PutUserHandler)

		// ユーザー退会機能
		user.DELETE("", DeleteUserHandler)
	}

	todo := r.Group("/todos")
	{
		todo.Use(MWGetUserID())

		// Todoリストを取得する
		todo.GET("", GetTodoHandler)

		// Todoを作成する
		todo.POST("", PostTodoHandler)
		todoPutDelete := todo.Group("/:todo_id")
		{
			// 指定されたTodoが認証されたユーザーが作成したものかチェックする
			todoPutDelete.Use(TodoCheckUser())

			// Todo情報を更新する
			todoPutDelete.PUT("", PutTodoHandler)

			// Todoを削除する
			todoPutDelete.DELETE("", DeleteTodoHandler)
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
