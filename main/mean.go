// 用于计算数组中数据的平均值
package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 4, 6, 7, 8, 13}
	res := mean(a)
	fmt.Println(res)
}

func mean(num []int) (ret float64) {
	var sum float64 = 0
	for _, value := range num {
		sum += float64(value) // 类型转换
	}
	return sum / float64(len(num))
}
