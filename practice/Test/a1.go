package Test

import "fmt"

type TZ int

func (tz *TZ) Increase(num int) {
	fmt.Println(*tz)
	*tz += TZ(num)
}

func TestTZ() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}
