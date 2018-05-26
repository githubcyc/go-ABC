// http://www.runoob.com/go/go-slice.html

// Go 语言切片是对数组的抽象。
// Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
// 与数组相比切片的长度是不固定的，可以追加元素，
// 在追加时可能使切片的容量增大。
// 1. 声明一个未指定大小的数组来定义切片：

// var identifier []type
// 2. 或使用make()函数来创建切片:

// var slice1 []type = make([]type, len)
// 也可以简写为

// slice1 := make([]type, len)
package main

import "fmt"

func main() {
	// make([]T, length, capacity)
	var numbers = make([]int, 3, 5)
	/* 向切片添加1 or 多个元素 */
	numbers = append(numbers, 1, 2)
	numbers[0] = 1
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	fmt.Println("number[0:2] ", numbers[0:2])
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
