package seed

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql","user1:pass@tcp(mysql:3306)/qin-todo")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Openエラー")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Pingエラー")
	}
	
}
