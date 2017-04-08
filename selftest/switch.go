package main

import (
	"fmt"
)

func main() {
	var (
		ch rune
	)
	fmt.Println("请输入成绩 大写字母 A ~ D")
	fmt.Scanf("%c", &ch)
	f(ch)
}

func f(x rune) {
	switch x {
	case 'A':
		fmt.Println("85 ~ 100")
	case 'B':
		fmt.Println("70 ~ 84")
	case 'C':
		fmt.Println("60 ~ 69")
	case 'D':
		fmt.Println("< 60")
	default:
		fmt.Println("error")
	}
}
