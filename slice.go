package main

import "fmt"

func main() {
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
	//cap()的使用
	mySlice1 := make([]int, 5, 10)
	fmt.Println("len(mySlice1):", len(mySlice1))
	fmt.Println("cap(mySlice1):", cap(mySlice1))
	//增加另一个数组切片
	mySlice2 := []int{7, 8, 9}
	mySlice = append(mySlice, mySlice2...)
	//省略号表示将mySlice2打散后传入
	for _, value := range mySlice {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")
	//基于数组切片创建数组切片
	newSlice := mySlice[5:]
	for _, value1 := range newSlice {
		fmt.Print(value1, " ")
	}
	fmt.Print("\n")
	//内容复制数组切片支持Go语言的另一个内置函数copy()，用于将内容从
	//一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，
	//就会按其中较小的那个数组切片的元素个数进行复制。
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	for _, var1 := range slice2 {
		fmt.Print(var1, " ")
		// 1 2 3
	}
	fmt.Print("\n")
	slice3 := []int{5, 4, 3}
	copy(slice1, slice3) // 只会复制slice3的3个元素到slice1的前3个位置
	for _, var2 := range slice1 {
		fmt.Print(var2, " ")
		//5 4 3 4 5
	}
}
