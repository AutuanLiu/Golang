/*
*  求 11-999 的回文数 使 m = m^2 =m^3
*
 */

package main

import (
	"fmt"
)

func main() {
	for m := 10; m < 1000; m++ {
		if sym(m) && sym(m*m) && sym(m*m*m) {
			fmt.Println("回文数 ", m)
		}
	}
}

func sym(m int) (b bool) {
	var (
		n int = 0
		i int = m
	)
	for i != 0 { //÷ 10 取余 构造新数
		n = n*10 + i%10
		i /= 10
		b = (m == n)
	}
	return
}
