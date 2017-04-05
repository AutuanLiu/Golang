package main

import (
	"fmt"
)

func main() {
	for i := 1896; i <= 1910; i++ {
		if isLeapYear(i) {
			fmt.Println(i, "是闰年")
		} else {
			fmt.Println(i, "不是闰年")
		}
	}
}

// 1900 刚好是一个特例
func isLeapYear(year int) bool {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		return true
	} else {
		return false
	}
}
