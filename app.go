package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Database Connectivity")
	db, err := sql.Open("mysql", "root@/golangdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully Connected to MySql Database")
	/// Insert the record in the table /
	insert, err := db.Query("INSERT INTO users VALUES ('','TEST' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
