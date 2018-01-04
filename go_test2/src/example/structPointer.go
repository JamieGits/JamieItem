package main
/**
 *  结构体指针
 */
import "fmt"

type  Books struct {
	title string
	author string
	subject string
	book_id int
}

func printBooks(books *Books)  {
	fmt.Printf("Book tile : %s\n",books.title)
	fmt.Printf("Book autor : %s\n",books.author)
	fmt.Printf("Book subject : %s\n",books.subject)
	fmt.Printf("Book book_id : %d\n",books.book_id)
	fmt.Println(*books)
	fmt.Printf("%+v\n",*books)

}
func main()  {
	var book1 Books
	var book2 Books

	book1.title = "《平凡的世界》"
	book1.author = "路遥"
	book1.subject = "平凡的世界"
	book1.book_id = 12

	book2.title = "《红楼梦》"
	book2.author = "曹雪芹"
	book2.subject = "红楼梦"
	book2.book_id = 13

	printBooks(&book1)
	fmt.Printf("\n\n")
	printBooks(&book2)

}