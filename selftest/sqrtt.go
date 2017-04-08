// 一元二次方程
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a,
		b,
		c float64
		flag = true
	)
	fmt.Println("请按顺序输入方程的各项系数：")
	for flag {
		if math.Abs(a) < 1e-6 {
			fmt.Scanf("%f%f%f", &a, &b, &c)
			fmt.Printf("方程为：%f x^2 + %f x + %f\n", a, b, c)
			fmt.Println("请重新输入 ")
		} else {
			flag = false
			delta := math.Pow(b, 2) - 4*a*c
			if delta > 0 {
				x1 := (-b + math.Sqrt(delta)) / (2 * a)
				x2 := (-b - math.Sqrt(delta)) / (2 * a)
				fmt.Println("结果是：两异根")
				fmt.Println("x1 = ", x1, "\n", "x2 = ", x2)
			} else if delta == 0 {
				ret := -b / (2 * a)
				fmt.Println("结果是：两同根")
				fmt.Println("x = ", ret)
			} else {
				// 复数类型的处理
				// x3 := (-b)/(2*a) + math.Sqrt(-delta)i
				// x4 := (-b)/(2*a) - math.Sqrt(-delta)i
				// 不对
				x3 := complex((-b)/(2*a), math.Sqrt(-delta))
				x4 := complex((-b)/(2*a), -math.Sqrt(-delta))
				fmt.Println("结果是：两复根")
				fmt.Println("x1 = ", x3, "\n", "x2 = ", x4)
			}
		}
	}
}
