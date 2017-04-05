package main

import "fmt"

func main() {
	var (
		a,
		b int
	)
	fmt.Scanf("%d%d", &a, &b)
	fsum(a, b)
}

func fsum(a, b int) {
	sum := float64(a) + float64(b)
	fmt.Println(sum)
}
