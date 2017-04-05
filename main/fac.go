// 实现阶乘的递归函数算法
package main

import "fmt"

const N int = 30

func main() {
	var flag bool = false
	for i := 0; i < N; i++ {
		if flag && i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("fac(%d)=%d,", i, fac(i))
		flag = true
	}
}

func fac(num int) (result int) {
	if num == 0 {
		result = 1
	} else {
		result = num * fac(num-1)
	}
	return result
}
