//不定参数类型
//不定参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要
//将函数定义为接受不定参数类型：
package main

import (
	"fmt"
)

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Print(arg, " ")
	}
}
func main() {
	myfunc(2, 3, 4)
	fmt.Print("\n")
	myfunc(1, 3, 7, 13)
}
