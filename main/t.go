package main

import "fmt"
import "reflect"

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func main() {
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
	var arr = [5]int{0, 1, 2, 3, 4}
	x := sum(arr[:]) // make a slice
	fmt.Println(x)
	tagf()
}

func tagf() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}
func f(a [3]int)   { fmt.Println(a) }
func fp(a *[3]int) { fmt.Println(a) }
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
