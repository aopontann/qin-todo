package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type TodoList struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Execution_date string `json:"execution_date"`
}

func InitRouter() *gin.Engine {
	// DB接続
	db, err := sql.Open("mysql", "user1:pass@tcp(mysql:3306)/qin-todo")
	if err != nil {
		log.Fatal(err)
	}
	// 一旦コメントアウトしておく
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// OAuth設定
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

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// todoリスト取得機能(デモ版)
	r.GET("/todo_list", func(c *gin.Context) {
		var (
			id             string
			content        string
			execution_date string
		)
		todo_list := []TodoList{}
		rows, err := db.Query("SELECT id, content, execution_date FROM todo_list")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &content, &execution_date)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, content, execution_date)
			todo_list = append(todo_list, TodoList{Id: id, Content: content, Execution_date: execution_date})
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		err = json.NewEncoder(os.Stdout).Encode(todo_list)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{
			"items": todo_list,
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

			// トークンを使ってgoogleアカウント情報を取得する
			resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + tok.AccessToken)
			if err != nil {
				log.Fatal(err)
			}
			// これやる意味わかってない...
			defer resp.Body.Close()

			var r io.Reader = resp.Body
			r = io.TeeReader(r, os.Stderr)

			var userInfo GoogleUserInfo
			err = json.NewDecoder(r).Decode(&userInfo)
			if err != nil {
				log.Fatal(err)
			}

			// ULIDの作成
			t := time.Now()
			entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
			id := ulid.MustNew(ulid.Timestamp(t), entropy)

			// googleアカウント情報をDBに保存する
			// id.String()とすることで正常に保存できるようになった。普通のidもulidを返しているが、
			// valuesにidを指定すると空白で保存される現象が発生した
			_, err = db.Exec("INSERT IGNORE INTO users (id, name, email, avatar_url, token) VALUES (?, ?, ?, ?, ?)", id.String(), userInfo.Name, userInfo.Email, userInfo.Picture, tok.AccessToken)
			if err != nil {
				log.Fatal(err)
			}

			c.JSON(200, gin.H{
				"ulid": id,
				"tok":  tok,
				"body": userInfo,
			})
		})
	}

	return r
}
