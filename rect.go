// 定义一个矩形类
// Go 中的面向对象使用 struct 实现
package main

import (
	"fmt"
)

type Rect struct {
	x, y          float64
	width, height float64
}

//定义成员方法Area()来计算矩形的面积
func (r *Rect) Area() float64 {
	return r.width * r.height
}
func main() {
	// 几种初始化方法与测试
	var r Rect
	r.height, r.width = 4, 3
	fmt.Println(r.Area())
	rect1 := new(Rect)
	fmt.Println(rect1.Area()) // 为显示赋值，则自动赋值 0
	rect2 := &Rect{}
	fmt.Println(rect2.Area())
	rect3 := &Rect{0, 0, 100, 200}
	fmt.Println(rect3.Area())
	rect4 := &Rect{width: 100, height: 200}
	fmt.Println(rect4.Area())
}
