// var identifier type
package main

import . "fmt"

// declare

//第一种，指定变量类型，声明后若不赋值，使用默认值。 e
// 第二种，根据值自行判定变量类型。 c, d
// 第三种，省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。d

// 全局变量是允许声明, 但不使用
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	j, l int
	k    bool
)
var x, y int // 0, 0
var a, b string = "tutorial", "go.com"
var c, d = 123, "hello"
var e bool
var g, h = 1, 2

// d := 10   X
// false

func main() {
	Println("d is ", d)
	d := 10 // in func
	//这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	println(x, y, a, b, c, d, g, h)
}
