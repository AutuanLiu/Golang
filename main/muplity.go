package main

import "fmt"
import "time"

var num int = 10
var numx2, numx3 int

// 为了程序的易读性，一般将 main函数放在前面，自定义的函数使用字母或者
// 某种自定的顺序排列起来
func main() {
	start := time.Now()
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
	fmt.Printf("The min is : %d \n", min(1, 5, -9, 1, 6))
	f()

	// 匿名函数
	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("\nmain took this amount of time: %s\n", delta)
}

func PrintValues() {
	fmt.Printf("num = %d, 2 times num = %d, 3 times num = %d\n", num, numx2, numx3)
}

// 未命名返回值
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名返回值
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// 变长参数的函数
func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

// defer function
func f() {
	for t := 0; t < 5; t++ {
		defer fmt.Printf("%d ", t)
	}
}
