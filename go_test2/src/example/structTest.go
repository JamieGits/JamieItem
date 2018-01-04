package main

import "fmt"
/**
 *  Go语言结构体访问
 */
type Books struct {
	title string
	author  string
	subject  string
	book_id   int
}
/* 结构体作为函数参数 */
func PrintBook(book Books)  {
	fmt.Printf("book1 title : %s\n",book.title)
	fmt.Printf("book1 author : %s\n",book.author)
	fmt.Printf("book1 subject : %s\n",book.subject)
	fmt.Printf("book1 book_id : %d\n",book.book_id)

}
func main()  {
	var book1  Books
	var book2  Books


	/* book1 描述*/
	book1.title = "Go语言"
	book1.author = "www.runoob.com"
	book1.subject = "Go语言教程"
	book1.book_id = 12345

	/* book2 描述*/
	book2.title = "Python 语言"
	book2.author = "www.runoob.com"
	book2.subject = "Python语言教程"
	book2.book_id = 23456


	/* 打印 book1的信息 */
	fmt.Printf("book1 title : %s\n",book1.title)
	fmt.Printf("book1 author : %s\n",book1.author)
	fmt.Printf("book1 subject : %s\n",book1.subject)
	fmt.Printf("book1 book_id : %d\n",book1.book_id)


   /* 打印 book2的信息 */
   fmt.Printf("book2 title : %s\n",book2.title)
   fmt.Printf("book2 author : %s\n",book2.author)
   fmt.Printf("book2 subject : %s\n",book2.subject)
   fmt.Printf("book2 boo_id : %d\n",book2.book_id)

   PrintBook(book1)
   PrintBook(book2)



}

