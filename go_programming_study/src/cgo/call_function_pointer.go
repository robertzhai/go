package main

import "fmt"

/*
#include<math.h>
typedef int (*intFunc) ();

 int proxy_int_func(intFunc f)
 {
		return f();
 }

 int userfunc()
 {
	    return 1;
 }
 */
import "C"

func main() {
	f := C.intFunc(C.userfunc)
	fmt.Println(int(C.proxy_int_func(f)))
	var n = C.sqrt(1)
	fmt.Print(n)
}