// 切片的练习
package main

import "fmt"

func main() {
	slice2()
}

func slice2() {
	s1 := []int{2, 3, 4, 5, 7, 13}
	s2 := make([]string, 5, 10)
	var s3 []string = make([]string, len(s2))
	m1 := make(map[int]string, 5)
	fmt.Println(s1)
	s1 = append(s1, 2, 4, 6, 7)
	fmt.Println(s1)
	s2 = []string{"l", "i", "u", "y", "a", "n"}
	copy(s3, s2) // 对齐原则
	fmt.Println(s3)
	m1[0] = "l"
	m1[1] = "i"
	m1[2] = "u"
	value, ok := m1[1]
	if ok {
		fmt.Println(value)
	}
	s4 := s1[4:5]
	fmt.Printf("%d, %d, %d", len(s4), cap(s1), cap(s4))
}
