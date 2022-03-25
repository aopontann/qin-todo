// このファイルにはお試しで実装してみるイベントハンドラーやサンプルを作成する
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aopontann/qin-todo/common"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	// "golang.org/x/crypto/bcrypt"
)

type TodoList struct {
	Id             string `json:"id"`
	Content        string `json:"content"`
	Execution_date string `json:"execution_date"`
}

type UserInfo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Avatar_url string `json:"avatar_url"`
}

type RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Pon(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// 全てのToDoを取得する機能
func GetTodoList(c *gin.Context) {
	// DB接続プールを取得する
	db := common.GetDB()
	var (
		id             string
		content        string
		execution_date string
	)
	todo_list := []TodoList{}
	// SQLを実行（全てのToDoからidと内容とやる日を出力する。）
	rows, err := db.Query("SELECT id, content, execution_date FROM todo_list")
	if err != nil {
		log.Fatal(err)
	}
	// これやる意味はわかってない
	// 調べた感じ「開いた結果セットがある限り、基礎となる接続はビジー状態であり、他のクエリに使用することはできない。つまり、コネクションプールで利用できない」と説明があり、ビジー状態を解除するためにClose()する必要があるのではないかと思う。
	defer rows.Close()
	// 行を反復処理する
	for rows.Next() {
		// カラムを変数に読み込む
		err := rows.Scan(&id, &content, &execution_date)
		if err != nil {
			log.Fatal(err)
		}
		// 確認でログ出力する
		log.Println(id, content, execution_date)
		// スライス(todo_list)に新しい要素を追加したスライスを作成しtodo_listに代入する
		todo_list = append(todo_list, TodoList{Id: id, Content: content, Execution_date: execution_date})
	}
	// 反復中に発生したエラーを返す。エラーが発生していない場合nilを返す
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// ステータスコード200でクライアントにtodo_listをレスポンスとして返す
	// gin.Hは map[string]interface{} のショートカット
	c.JSON(200, gin.H{
		"items": todo_list,
	})
}

// ユーザー情報取得機能(ハードコーディング版)
func GetUserHardCode(c *gin.Context) {
	userInfo := UserInfo{Id: "id001", Name: "taro", Avatar_url: "https://lh3.googleusercontent.com/a-/AOh14Gg7m3sGmgDctni57nyWg6ATJLrSJNeT4mKIPtb_lxo=s96-C"}

	c.JSON(200, userInfo)
}

// request body からメールとパスワードを取得してDBに保存する機能
func PostUserDemo(c *gin.Context) {
	var reqb RequestBody
	// byte(unit8)型の長さが2048のスライスを作成
	buf := make([]byte, 2048)
	// リクエストボディの内容をbufの先頭から埋めていく感じだと思う
	// 埋まったバイト数と、埋めていく過程で発生したエラーを返す
	n, _ := c.Request.Body.Read(buf)
	// byte型スライスからGo構造体にデコード
	err := json.Unmarshal(buf[:n], &reqb)
	if err != nil {
		log.Fatal(err)
	}
	// ULIDの作成
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	// `GenerateFromPassword` でパスワードをハッシュ化する。
	// 第2引数はコストを指定する。値は 4 ~ 31 の範囲である必要がある。
	// コストについてはこの記事がわかりやすい。 → https://qiita.com/istsh/items/ca330d27fe51a6bf7a3d#2-%E3%82%B3%E3%82%B9%E3%83%88%E3%81%AE%E6%8C%87%E5%AE%9A
	// hashed, _ := bcrypt.GenerateFromPassword([]byte(reqb.Password), 10)
	// DB接続プールを取得する
	db := common.GetDB()
	// paswordはハッシュ化してDBに保存した方がいいが、今はそのまま保存するようにしておく。
	_, err = db.Exec("INSERT INTO users (id, name, email, password) VALUES (?,?,?,?)", id.String(), "名前", reqb.Email, reqb.Password)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(201, gin.H{
		"reqb": reqb,
	})
}

// Cookieをいじってみる
func CookieDemo(c *gin.Context) {
	var json RequestBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test!", 3600, "/", "localhost", false, true)
	}

	fmt.Printf("Cookie  value: %s \n", cookie)
	c.JSON(200, json)
}
