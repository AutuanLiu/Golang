// 求最大公约数
package main

import (
	"fmt"
)

func main() {
	var (
		a,
		b int
	)
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println("递归调用 ", gcd(a, b))
	fmt.Println("非递归调用 ", gcdn(a, b))
	fmt.Println("最小公倍数 ", gcm(a, b))
}

// 递归形式
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// 非递归形式
func gcdn(a, b int) int {
	for b != 0 {
		t := a % b
		a = b
		b = t
	}
	return a
}

// 最小公倍数
func gcm(a, b int) int {
	return a * b / gcdn(a, b)
}
