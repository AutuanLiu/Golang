// 用于字符的测试
package main

import (
	"fmt"
	"reflect"
)

func main() {
	c1 := '我'
	var c2 rune = '是'
	fmt.Println(c1)
	fmt.Printf("%q%q\n", c1, c2)
	fmt.Printf("%b \n", c1)
	rtype := reflect.TypeOf(c2).Kind()
	fmt.Println(rtype)
}
