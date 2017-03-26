//用interface{}传递任意类型数据
package main

import "fmt"

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case float32:
			fmt.Println(arg, "is a float32 value.")
		case complex64:
			fmt.Println(arg, "is a complex64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	var v5 float64 = 1.356
	var v6 complex64 = 1.6 + 5.3i
	MyPrintf(v1, v2, v3, v4, v5, v6)
}
