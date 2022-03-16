package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("SECRET_KEY"),
		RedirectURL:  "http://localhost:18080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		// google認証画面にリダイレクト
		auth.GET("/google", func(c *gin.Context) {
			url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
			c.Redirect(http.StatusMovedPermanently, url)
		})
		// トークン取得エンドポイント
		auth.GET("/token", func(c *gin.Context) {
			code := c.Query("code")
			// 認可コードからトークンを取得する
			tok, err := conf.Exchange(oauth2.NoContext, code)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(200, tok)
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
