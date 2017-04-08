package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(f(a))
}

func f(x int) int {
	if x < 0 {
		return -1
	} else if x == 0 {
		return 0
	} else {
		return 1
	}
}
