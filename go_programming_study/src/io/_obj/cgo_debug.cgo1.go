// Created by cgo - DO NOT EDIT

//line /Users/baidu/gotrain/src/io/cgo_debug.go:1
package main
//line /Users/baidu/gotrain/src/io/cgo_debug.go:7

//line /Users/baidu/gotrain/src/io/cgo_debug.go:6
import (
	"unsafe"
	"fmt"
)
//line /Users/baidu/gotrain/src/io/cgo_debug.go:12

//line /Users/baidu/gotrain/src/io/cgo_debug.go:11
func Print(s string) {
	cs := _Cfunc_CString(s)
	_Cfunc_fputs(cs, (*_Ctype_struct___sFILE)(*_Cvar_stdout))
	_Cfunc_free(unsafe.Pointer(cs))
}
//line /Users/baidu/gotrain/src/io/cgo_debug.go:18

//line /Users/baidu/gotrain/src/io/cgo_debug.go:17
func main() {
	fmt.Print("cgo")
	Print("test cgo")
}
