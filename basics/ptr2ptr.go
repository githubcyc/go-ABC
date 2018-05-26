// 如果一个指针变量存放的又是另一个指针变量的地址，
// 则称这个指针变量为指向指针的指针变量
package main

import "fmt"

func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	// 访问指向指针的指针变量值需要使用两个 * 号
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
