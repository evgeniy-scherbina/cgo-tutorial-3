package main

// #include "hello.h"
import "C"

func main() {
	name := C.CString("Gopher")
	C.hello(name)
}
