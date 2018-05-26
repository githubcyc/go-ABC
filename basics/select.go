/*
select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
    //你可以定义任意数量的 case
    default : // 可选
       statement(s);
}
*/
// 每个case必须是一个通信操作，要么是发送要么是接收。
// 所有channel表达式都会被求值
package main

import "fmt"

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
