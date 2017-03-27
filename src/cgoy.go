package main

import "fmt"

/*
#include <stdio.h>
void hello() {
printf("Hello, Cgo! -- From C world.\n");
}
*/
import "C"

func Hello() int {
	return int(C.hello())
}

func main() {
	fmt.Println(Hello())
}
