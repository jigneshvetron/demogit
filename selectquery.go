package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// database Open connection
	db, err := sql.Open("mysql", "root@/golangdb")
	if err != nil {
		panic(err.Error())
	}
	// main function close
	defer db.Close()
	select1, err := db.Query("select username from users")
	if err != nil {
		panic(err.Error())
	}
	type Test struct {
		title string `json:"title"`
	}
	for select1.Next() {
		var test Test
		err = select1.Scan(&test.title)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Println(test.title)
	}
}
