// 题目来自于 c 程序设计 谭浩强
package main

import (
	"fmt"
)

const N int = 30

func main() {
	pstars()
	pchar()
	pstars()
}

func pchar() {
	fmt.Println()
	fmt.Println("          Very Good!")
}

func pstars() {
	for i := 0; i < N; i++ {
		fmt.Print("*")
	}
}
