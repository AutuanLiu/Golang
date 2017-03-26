//Go语言中的闭包
package main

import (
	"fmt"
)

func main() {
	var j int = 5
	a := func() func() { //返回一个函数
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	// 在变量a指向的闭包函数中，只有内部的匿名函数才能访问变量i，而无法通过
	// 其他途径访问到，因此保证了i的安全性
	a()
	j *= 2
	a()
}
