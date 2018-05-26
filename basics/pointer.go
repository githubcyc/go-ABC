// 变量是一种使用方便的占位符，用于引用计算机内存地址。

// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

// 以下实例演示了变量在内存中地址

package main

import "fmt"

func main() {
	var a int = 10 /* 声明实际变量 */

	fmt.Printf("address: %x\n", &a)

	var ip *int /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	// 空指针
	var ptr *int
	/* ptr 不是空指针 */
	if ptr != nil {
		fmt.Println("not nil")
	}
	/* ptr 是空指针 */
	if ptr == nil {
		fmt.Println("ptr is nil")
	}
	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

// 一个指针变量指向了一个值的内存地址。

// 类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
// var var_name *var-type
// var ip *int        /* 指向整型*/
// var fp *float32    /* 指向浮点型 */
