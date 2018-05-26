// Go 语言中数组可以存储同一类型的数据，
// 但在结构体中我们可以为不同项定义不同的数据类型。

// 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go"
	Book1.author = "www.google.com"
	Book1.subject = "Go tutorial"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python tutorial"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)
}
