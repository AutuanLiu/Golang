// 求素数
package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if isPrime(a) {
		fmt.Println(a, "是素数")
	} else {
		fmt.Println(a, "不是素数")
	}
}

func isPrime(num int) (b bool) {
	if num < 2 {
		fmt.Println("输入的数据不符合条件")
		os.Exit(0)
	} else if num == 2 {
		b = true
	} else {
		i := 2
		// GoLang 不支持隐式类型转换
		bound := int(math.Sqrt(float64(num))) + 1
		// 提前计算边界，宁多不少
		for i <= bound {
			if num%i == 0 {
				b = false
				break
			} else {
				b = true
			}
			i++
		}
	}
	return
}
