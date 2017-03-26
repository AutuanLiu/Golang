// 新类型Integer，它和int没有本质不同，只是它为内置的int类型增加了个新方法Less()。
// 这样实现了Integer后，就可以让整型像一个普通的类一样使用
package main

import (
	"fmt"
)

type Integer int

// 面向对象方法
func (a Integer) Less(b Integer) bool {
	return a < b
}

// 非面向对象方法
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}
func main() {
	var a Integer = 1
	if a.Less(2) { // 区别
		fmt.Println(a, "Less 2")
	} else {
		fmt.Println(a, "bigger than 2")
	}
	if Integer_Less(a, 2) { // 区别
		fmt.Println(a, "smaller than 2")
	} else {
		fmt.Println(a, "bigger than 2")
	}
}
