package main

import (
	"./src/Cal"
	"./src/Cal/Test"
	"fmt"
)

func main() {
	a, b := 2, 1
	fmt.Println(Cal.Add(a, b))
	fmt.Println(Cal.Sub(a, b))
	fmt.Println(Test.Abs(-1))
	fmt.Println(Cal.Multi(a, b))
}
