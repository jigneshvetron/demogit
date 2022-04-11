package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root@/golangdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	type Test struct {
		title string `json:"title"`
	}

	if err != nil {
		panic(err.Error())
	}
	var test Test
	err1 := db.QueryRow("SELECT username FROM users where id = ?", 2).Scan(&test.title)
	if err1 != nil {
		panic(err1.Error())
	}
	log.Println(test.title)

}
