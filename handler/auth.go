package handler

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"golang.org/x/oauth2"
	"github.com/aopontann/qin-todo/common"
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
