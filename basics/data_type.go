// http://www.runoob.com/go/go-data-types.html
package main

import "fmt"

var isActive bool                   // 全局变量声明
var enabled, disabled = true, false // 忽略类型的声明
func test() {
	var available bool // 一般声明
	valid := false     // 简短声明
	// available = true   // 赋值操作
	fmt.Println(available, valid)
}

func main() {
	var a = 1.5
	var b = 2
	test()
	fmt.Println(a, b)
}
