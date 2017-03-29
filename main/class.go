// 类的例子
package main

import "fmt"

// 相当于成员变量
type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

// 绑定在某类型的成员方法
func (tn *TwoInts) AddThem() (c int) {
	c = tn.a + tn.b
	return
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
