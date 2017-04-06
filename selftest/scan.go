package main

import (
	"fmt"
)

func main() {
	f1()
	f2()
	f3()
	f4()
	f5()
}

// 对于 Scan 而言，回车视为空白
func f1() {
	a, b, c := "", 0, false
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 abc 1 回车 t 回车
	// 结果 abc 1 true
}

// 整型变量也可以解析八进制和十六进制
func f2() {
	a, b, c := "", 0, false
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 def 0x20 T 回车
	// 结果 def 32 true
}

// 格式字符串可以指定宽度
func f3() {
	a, b, c := "", 0, false
	fmt.Scanf("%4s%d%t", &a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 1234567true 回车
	// 结果 1234 567 true
}

// 指定宽度不能跨越空白
func f4() {
	a, b, c := "", 0, false
	fmt.Scanf("%4s%d%t", &a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 12 34567True 回车
	// 结果 12 34567 true
}

// %c 总是匹配下一个字符，包括空格
func f5() {
	a, b, c := "", 0, 0
	fmt.Scanf("%s%c%d", &a, &b, &c)
	fmt.Println(a, b, c)
	// 在终端执行后，输入 abc 1 回车
	// 结果 abc 32 1
}
