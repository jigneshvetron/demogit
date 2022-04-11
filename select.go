package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Mysql Tutorial")
	// Open database Connection
	db, err := sql.Open("mysql", "root@/golangdb")
	// this iS error Open handling Error
	if err != nil {
		log.Print(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// This Is select Query
	select1, err := db.Query("select * from users order by id desc")

	if err != nil {	
		panic(err.Error())
	}
	type Tag struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	for select1.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = select1.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Println(tag.ID)
		log.Printf(tag.Name)
		defer select1.Close()
	}

}
