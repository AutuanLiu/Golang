package main

import (
	"fmt"
)

type Integer int

//面向对象方法
func (a *Integer) add(b Integer) { // 包内方法
	*a += b
}

// 只是在方法内修改了内容
func (a Integer) add1(b Integer) {
	a += b
	fmt.Println(a)
	// 方法内输出 a 的值
}

func (a Integer) add2(b Integer) Integer {
	return a + b
}
func main() {
	var a Integer = 1
	//调用add1()方法
	a.add(3)
	fmt.Println(a) // 4
	//调用add2()方法
	fmt.Println(a.add2(3)) // 7
	a.add1(5)              // 9
}
