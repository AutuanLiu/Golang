//实现在Go中调用C语言标准库的puts函数
package main

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("Hello, world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
