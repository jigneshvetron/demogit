package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go My Sql Tutorial")
	db, err := sql.Open("mysql", "root@/golangdb")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Databse Connect Successfully")
	insert, err := db.Query("insert into users values('','Good Morning')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
