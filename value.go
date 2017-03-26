//值类型和引用
//Go语言中的数组和基本类型没有区别，是很纯粹的值类型
package main

import (
	"fmt"
)

func main() {
	var a = [3]int{1, 2, 3}
	var b = a
	// b=a赋值语句是数组内容的完整复制
	b[1]++
	fmt.Println(a, b)
	var c = [3]int{1, 2, 3}
	//要想表达引用，需要用指针
	var d = &c
	d[1]++
	fmt.Println(c, *d)
}
