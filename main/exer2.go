package main

import (
	"fmt"
)

func main() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	// goto 的使用
	a := 1
	b := 9
	goto TARGET // compile error
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
