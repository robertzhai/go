package main

/*
#include <stdio.h> 
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func main() {
	cstr := C.CString("Hello cgo")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}