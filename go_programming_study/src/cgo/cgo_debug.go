package main
/*
#include <stdio.h>
#include <stdlib.h>
void output(char *str) {
    printf("%s\n", str);
}
*/
import "C" //此处和上面的C代码一定不能有空行 否则会报错
import "unsafe"
func main() {
	str := "hello cgo"
	//change to char*
	cstr := C.CString(str)
	C.output(cstr)
	C.free(unsafe.Pointer(cstr))
}