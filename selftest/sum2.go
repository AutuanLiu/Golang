package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sum2())
}

func sum2() (sum1 float64) {
	var sign float64 = -1
	var t float64
	for i := 1; i <= 100; i++ {
		t = float64(i)
		sum1 += math.Pow(sign, t+1.0) * 1.0 / t
	}
	return sum1
}
