package common

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Conf *oauth2.Config

// Google認証に必要な設定
// main.goで.envの読み込みをしてこの関数を実行しているが、関数外でconfの設定をしたらどうなるか
func GoogleAuthInit() {
	Conf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("SECRET_KEY"),
		RedirectURL:  "http://localhost:18080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
