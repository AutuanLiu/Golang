/*
*    二分幂法 求x^n
 */

// 求整数幂
package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var n int
	fmt.Scanf("%f%d", &x, &n)
	fmt.Println(powerf(x, n))
	fmt.Println(powerf2(x, n))
	fmt.Println(powerf3(x, n))
	fmt.Println(math.Pow(x, float64(n)))
}

func powerf(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}

/*
*    递归法 求x^n
 */
func powerf2(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else {
		return x * powerf2(x, n-1)
	}
}

/*
*    循环法 求x^n
 */

func powerf3(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}
