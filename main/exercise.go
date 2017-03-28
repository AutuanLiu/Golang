//个人案例测试
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%04d年%02d月%02d日\n", t.Year(), t.Month(), t.Day())
	fmt.Println(t)
	fmt.Printf("%02d点%02d分%02d秒", t.Hour(), t.Minute(), t.Second())
	s := "abcdefg"
	fmt.Printf("%s HasPrefix\n", s)
	ret := strings.HasPrefix(s, "abc")
	fmt.Printf("%t", ret)
	fmt.Println(strings.Contains(s, "cd"))
	fmt.Println()
	fmt.Println(strings.Index(s, "ef"))
	fmt.Println(strings.Replace(s, "a", "q", 0))
	fmt.Print(`help:
		1* p
		2* q`)
	fmt.Println()
	e := 1
	fmt.Println(&e)
	f := "good bye"
	var p *string = &f
	*p = "ping"
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", f)
	fmt.Println(math.Abs(-8))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Sin(math.Pi / 2))
	a := "123"
	c := "edf"
	b, _ := strconv.Atoi(a)
	m, ok := strconv.Atoi(c)
	fmt.Printf("%s to integer is %d", a, b)
	fmt.Println()
	if ok == nil {
		fmt.Printf("%s to integer is %d", c, m)
	} else {
		fmt.Println(c, "is not a Integer")
	}
	var num1 int = 100
	// 带条件
	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	var num2 int = 7
	// 不带条件
	switch {
	case num2 < 0:
		fmt.Println("Number is negative")
	case num2 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}

	str2 := "日本語中国"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	for index, rune := range str2 {
		fmt.Printf("%-2d %c\n", index, rune)
	}

	str := "G"
	for i := 1; i <= 5; i++ {
		println(str)
		str += "G"
	}
}
