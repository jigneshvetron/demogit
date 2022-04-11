package main
import "fmt"
type Books struct {
	title   string
	number  int
	subject string
	book_id int
}

func main() {
	var Book1 Books
	Book1.title = "This Is First title"
	Book1.number = 12
	// fmt.Println(Book1.title)
	// fmt.Println(Book1.number)
	structure(&Book1)


}
func structure(Books2 *Books){
  fmt.Println(Books2.title)
	fmt.Println(Books2.number)
}
