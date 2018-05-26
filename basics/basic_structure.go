// 当前程序的包名
package main

// 为fmt起别名为fmt2
// import fmt2 "fmt"
// import {
//     "fmt",
//     "io"
// }
// 导入其他包
// 前面加个点表示省略调用 == from A import *
import . "fmt"

// 通过 const 关键字来进行常量的定义。

// 通过在函数体外部使用 var 关键字来进行全局(global)变量的声明和赋值。

// 通过 type 关键字来进行结构(struct)和接口(interface)的声明。

// 通过 func 关键字来进行函数的声明。
// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main函数作为程序入口点启动
func main() {
	Println("Hello World!")
}

// 函数名首字母小写即为 private :

// func getId() {}
// 函数名首字母大写即为 public :

// func Printf() {}
