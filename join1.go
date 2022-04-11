package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// database connection
    fmt.Println("Mysql Tutorial")
	db, err := sql.Open("mysql", "root@/golangdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	selectdata, err := db.Query("SELECT users.id,employee.name FROM users INNER JOIN employee ON employee.user_id=users.id ")
	for selectdata.Next() {
		type Tag struct {
			ID   int    `json:"id1"`
			Name string `json:"namedsddd"`
		}
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = selectdata.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Println(tag.ID)
		log.Printf(tag.Name)
		defer selectdata.Close()
	}

}
