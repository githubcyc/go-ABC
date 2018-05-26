// 1.函数内定义的变量称为局部变量
// 2.函数外定义的变量称为全局变量
// 3.函数定义中的变量称为形式参数
// Go 语言程序中全局变量与局部变量名称可以相同，
// 但是函数内的局部变量会被优先考虑。
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
	/* 声明局部变量 */
	var a, b int

	/* 初始化参数 */
	a = 10
	b = 20
	var g int = 10
	var c int = 0

	fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
	c = sum(a, b)
	fmt.Printf("main()函数中 c = %d\n", c)
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}
