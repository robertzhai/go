package main 

/*
#include<stdio.h>

static void sayCgo(const char *s) {
	puts(s);
}

*/
import "C"

func main() {

	C.puts(C.CString("hello cgo\n"))
	C.sayCgo(C.CString("hello cgo\n"))

}