// 题目来自于 c 程序设计 谭浩强
package main

import "fmt"

func main() {
	var (
		a,
		b,
		c int
	)
	fmt.Scanf("%d%d%d", &a, &b, &c)
	m := fmax(a, b, c)
	fmt.Println(m)
}

func fmax(a, b, c int) int {
	max := a
	if max < b {
		max = b // 找到了 a ，b 中的最大值
	}
	if max < c {
		max = c
	}
	return max
}
