package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
	_, err = tx.Exec("INSERT INTO users(id, name, email, password) VALUES('1','太郎','abc@example.com','pass1'), ('2','二郎','def@example.com','pass2'), ('3','alex','ghi@example.com','pass3')")
	if err != nil {
		log.Fatal(err)
	}
	// デモtodoデータを追加
	_, err = tx.Exec("INSERT INTO todo_list(id, content, completed, execution_date, user_id) VALUES('1', 'デモ1', false, '2019-10-04 00:00:00', '1'), ('2', 'デモ2', true, '2019-10-05 00:00:00', '1'), ('3', 'デモ3', false, '2019-10-06 00:00:00', '2')")
	if err != nil {
		log.Fatal(err)
	}
	// コミット
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
