package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("My Sql Tutorial")
	// Open up our database connection.
	db, err := sql.Open("mysql", "root@/golangdb")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
    // executing
	defer db.Close()
}
