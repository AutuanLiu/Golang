// 闭包的测试函数
package main

import (
	"fmt"
)

func main() {
	c1 := f(1)
	c2 := f(2)
	x := c1()
	y := c2()
	fmt.Println(x, "  ", y)
}

func f(i int) func() int {
	i++
	return func() int { // 此处必须是匿名的
		i++
		return i
	}
}
