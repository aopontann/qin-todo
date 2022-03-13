package main

import (
	"log"
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
	r := gin.Default()
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth", func(c *gin.Context) {
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		c.JSON(200, gin.H{
			"url": url,
		})
	})

	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		tok, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Fatal(err)
		}
		client := conf.Client(oauth2.NoContext, tok)
		c.JSON(200, gin.H{
			"token": tok,
			"client": client.Get("...")
		})
	})

	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
