// 定义指向结构体的指针类似于其他指针变量，格式如下：

// var struct_pointer *Books

package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* Declare Book1 of type Book */
	var Book2 Books /* Declare Book2 of type Book */

	/* book 1 描述 */
	Book1.title = "Go"
	Book1.author = "www.Go.com"
	Book1.subject = "Go"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python"
	Book2.author = "www.Python.com"
	Book2.subject = "Python"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)
}
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
