// const identifier [type] = value
// 显式类型定义： const b string = "abc"
// 隐式类型定义： const b = "abc"

package main

import "fmt"

// import "unsafe"

// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
const (
	a = iota
	b
	c
)

// const (
// 	a = "abc"
// 	b = len(a)
// 	c = unsafe.Sizeof(a)
// )

func main() {
	println(a, b, c)
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("S : %d", area)
	println()
	println(a, b, c)
}
