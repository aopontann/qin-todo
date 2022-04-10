package handler

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/aopontann/qin-todo/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
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

func GoogleAuthRedirect(c *gin.Context) {
	url := common.Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusMovedPermanently, url)
}

func GoogleAuthGetToken(c *gin.Context) {
	code := c.Query("code")
	// 認可コードからトークンを取得する
	tok, err := common.Conf.Exchange(oauth2.NoContext, code)
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
	db := common.GetDB()
	_, err = db.Exec("INSERT IGNORE INTO users (id, name, email, avatar_url, token) VALUES (?, ?, ?, ?, ?)", id.String(), userInfo.Name, userInfo.Email, userInfo.Picture, tok.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"ulid": id,
		"tok":  tok,
		"body": userInfo,
	})
}

// リクエストボディからメールアドレスとパスワードを取得し、セッションIDを発行する
func SessionAuthLogin(c *gin.Context) {
	var (
		id   string
		pass string
	)
	var reqb RequestBody
	// リクエストボディを構造体にシリアライズする(Ginの機能)
	err := c.ShouldBindJSON(&reqb)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := common.GetDB()
	// emailが一致するユーザーのidとパスワードを取得する
	err = db.QueryRow("SELECT id, password FROM users WHERE email = ?", reqb.Email).Scan(&id, &pass)
	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong email address or password"})
		return
	}

	// パスワードを検証する
	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(reqb.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Wrong email address or password"})
		return
	}

	// uuidを生成する(セッションIDの生成)
	sid := uuid.NewString()

	// Cookieの"session"というkeyにuuidを保存する(有効期限はとりあえず1時間)
	c.SetCookie("session", sid, 3600, "/", "localhost", false, true)

	// redisにuuidをキーとしてユーザーidを値として保存する(これも有効期限を1時間とする)
	rdb := common.GetRDB()
	err = rdb.Set(c, sid, id, 1*time.Hour).Err()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":        id,
		"sessionId": sid,
		"reqb":      reqb,
	})
}

func SessionAuthLogout(c *gin.Context) {
	// cookieからセッションIDを取得
	sid, err := c.Cookie("session")
	if err != nil {
		c.JSON(401, gin.H{"error": "Already logged out. Or your session ID is invalid."})
		return
	}

	// maxAgeに0かマイナス値を指定することで対象のcookieを削除することができる
	c.SetCookie("session", "", -1, "/", "localhost", false, true)

	// redisに保存されているセッションIDを削除する
	rdb := common.GetRDB()
	if err := rdb.Del(c, sid).Err(); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Logged out.",
	})
}

// ユーザー登録機能
func UserRegister(c *gin.Context) {
	var reqb RequestBody
	// リクエストボディを構造体にシリアライズする(Ginの機能)
	if err := c.ShouldBindJSON(&reqb); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := common.GetULID()

	// パスワードをハッシュ化
	hashed, _ := bcrypt.GenerateFromPassword([]byte(reqb.Password), 10)

	// DB接続プールを取得する
	db := common.GetDB()
	// idと"名前"、メールアドレス、ハッシュ化したパスワードをDBに保存する(ユーザー名のデフォルトをとりあえず"名前"しておく)
	// 名前の決め方や何かいい案があれば教えてほしい
	_, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES (?,?,?,?)", id, "名前", reqb.Email, string(hashed))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"id":     id,
		"hashed": string(hashed),
		"reqb":   reqb,
	})
}
