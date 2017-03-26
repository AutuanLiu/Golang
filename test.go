package main

import (
	"fmt"
)

var num1, num2 int
var c = complex(3.2, 12)

func main() {
	num1 = 7
	num2 = 3
	var str string   // 声明一个字符串变量
	str = "Hello,世界" // 字符串赋值
	b := (num1 == num2)
	ch := str[0] // 取字符串的第一个字符
	fmt.Println("123" + "世纪长河")
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
	fmt.Println(num1+num2, b, num1%num2, ^6)
	fmt.Println(real(c), imag(c))
	if num1 == num2 {
		fmt.Println("num1 are equal to num2")
	} else {
		fmt.Println("num1 are not equal to num2")
	}
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}
}
