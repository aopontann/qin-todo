package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserSample struct {
	Id       string
	Name     string
	Email    string
	Password string
}

func main() {
	db, err := sql.Open("mysql", "user1:pass@tcp(mysql:3306)/qin-todo")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Openエラー")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Pingエラー")
	}

	// デモユーザー情報の作成
	var users []UserSample
	// パスワード"test123"で登録した時のユーザー情報(太郎だけアイコンがあるようにしてみる)
	users = append(users, UserSample{Id: "01FZ4NX3QG2FEMXS984JHEDV2B", Name: "太郎", Email: "test1@example.com", Password: "$2a$10$C4O/7yUOzokQVXNDd6pzPe4pyPMJ1Py4aIJ1HjG2Ed1y4qwPQ1e5."})
	// test456
	users = append(users, UserSample{Id: "01FZ4NY6195YHN33X736978NTV", Name: "二郎", Email: "test2@example.com", Password: "$2a$10$16IhNWV3s8BWHuLXoGmYJOD6xwsaXuJQWErLWqPAri1jUnEpUqVIK"})
	// test789
	users = append(users, UserSample{Id: "01FZ4NYFM8F2KYZP6GZ6Q15YW7", Name: "alex", Email: "test3@example.com", Password: "$2a$10$3knQjJFUPDegKEGbB/q3Nu6zFCLQ0gdJejzs0qA4QAlirUwDyaDTK"})

	// トランザクション
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	// todoテーブルの全データ削除
	_, err = tx.Exec("DELETE FROM todo_list")
	if err != nil {
		log.Fatal(err)
	}
	// usersテーブルの全データ削除 TRUNCATE TABLE users 使えんかった
	_, err = tx.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
	// デモユーザーデータを追加
	_, err = tx.Exec("INSERT INTO users(id, name, email, password, avatar_url) VALUES (?,?,?,?,?), (?,?,?,?,?), (?,?,?,?,?)", users[0].Id, users[0].Name, users[0].Email, users[0].Password, "https://lh3.googleusercontent.com/a-/AOh14Gg7m3sGmgDctni57nyWg6ATJLrSJNeT4mKIPtb_lxo=s96-c", users[1].Id, users[1].Name, users[1].Email, users[1].Password, nil, users[2].Id, users[2].Name, users[2].Email, users[2].Password, nil)
	if err != nil {
		log.Fatal(err)
	}
	// デモtodoデータを追加
	_, err = tx.Exec("INSERT INTO todo_list(id, content, completed, execution_date, user_id) VALUES('1', 'デモ1', false, '2019-10-04 00:00:00', '01FZ4NX3QG2FEMXS984JHEDV2B'), ('2', 'デモ2', true, '2019-10-05 00:00:00', '01FZ4NX3QG2FEMXS984JHEDV2B'), ('3', 'デモ3', false, '2019-10-06 00:00:00', '01FZ4NY6195YHN33X736978NTV')")
	if err != nil {
		log.Fatal(err)
	}
	// コミット
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
