package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("hello")
     fmt.Println("world")
}
