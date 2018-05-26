// 声明形参为数组，我们可以通过以下两种方式来声明：
// 1.  myFunction(param [10]int)
// 2.  myFunction(param []int)

package main

import "fmt"

func main() {
	/* 数组长度为 5 */
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, len(balance))

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
