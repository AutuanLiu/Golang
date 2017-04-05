// exp(x) 的估计
package main

import (
	"fmt"
	"math"
)

func expf(x float64) float64 {
	var (
		sum  float64 = 1.0 + x
		item float64 = x * x / 2.0
		i    int     = 3
	)

	for math.Abs(item) > 1e-7 {
		sum += item
		item = item * x / float64(i)
		i++
	}
	return sum
}

func main() {
	var x float64
	fmt.Scanf("%f", &x)
	fmt.Println("exp(", x, ")的近似值 ", expf(x))
	fmt.Println("exp(", x, ")的验证值 ", math.Exp(x))
	fmt.Println("误差 ", math.Abs(expf(x)-math.Exp(x)))
}
