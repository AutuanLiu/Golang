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
	for i := 1; i <= 100; i++ {
		sum1 += math.Pow(sign, float64(i+1)) * float64(1) / float64(i)
	}
	return sum1
}
