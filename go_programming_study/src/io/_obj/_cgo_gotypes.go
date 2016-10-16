// Created by cgo - DO NOT EDIT

package main

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_FILE _Ctype_struct___sFILE

type _Ctype___darwin_off_t _Ctype___int64_t

type _Ctype___int64_t _Ctype_longlong

type _Ctype_char int8

type _Ctype_fpos_t _Ctype___darwin_off_t

type _Ctype_int int32

type _Ctype_longlong int64

type _Ctype_short int16

type _Ctype_struct___sFILE struct {
//line :1
	_p		*_Ctype_unsignedchar
//line :1
	_r		_Ctype_int
//line :1
	_w		_Ctype_int
//line :1
	_flags		_Ctype_short
//line :1
	_file		_Ctype_short
//line :1
	_		[4]byte
//line :1
	_bf		_Ctype_struct___sbuf
//line :1
	_lbfsize	_Ctype_int
//line :1
	_		[4]byte
//line :1
	_cookie		unsafe.Pointer
//line :1
	_close		*[0]byte
//line :1
	_read		*[0]byte
//line :1
	_seek		*[0]byte
//line :1
	_write		*[0]byte
//line :1
	_ub		_Ctype_struct___sbuf
//line :1
	_extra		*_Ctype_struct___sFILEX
//line :1
	_ur		_Ctype_int
//line :1
	_ubuf		[3]_Ctype_unsignedchar
//line :1
	_nbuf		[1]_Ctype_unsignedchar
//line :1
	_lb		_Ctype_struct___sbuf
//line :1
	_blksize	_Ctype_int
//line :1
	_		[4]byte
//line :1
	_offset		_Ctype_fpos_t
//line :1
}

type _Ctype_struct___sFILEX struct{}

type _Ctype_struct___sbuf struct {
//line :1
	_base	*_Ctype_unsignedchar
//line :1
	_size	_Ctype_int
//line :1
	_	[4]byte
//line :1
}

type _Ctype_unsignedchar uint8

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cmalloc runtime.cmalloc
func _cgo_runtime_cmalloc(uintptr) unsafe.Pointer

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr)
//go:linkname __cgo___stdoutp __stdoutp
//go:cgo_import_static __stdoutp
var __cgo___stdoutp byte
var _Cvar_stdout **_Ctype_struct___sFILE = (**_Ctype_struct___sFILE)(unsafe.Pointer(&__cgo___stdoutp))


func _Cfunc_CString(s string) *_Ctype_char {
	p := _cgo_runtime_cmalloc(uintptr(len(s)+1))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*_Ctype_char)(p)
}
//go:cgo_import_static _cgo_4dbd5e71a156_Cfunc_fputs
//go:linkname __cgofn__cgo_4dbd5e71a156_Cfunc_fputs _cgo_4dbd5e71a156_Cfunc_fputs
var __cgofn__cgo_4dbd5e71a156_Cfunc_fputs byte
var _cgo_4dbd5e71a156_Cfunc_fputs = unsafe.Pointer(&__cgofn__cgo_4dbd5e71a156_Cfunc_fputs)

func _Cfunc_fputs(p0 *_Ctype_char, p1 *_Ctype_struct___sFILE) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_4dbd5e71a156_Cfunc_fputs, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_4dbd5e71a156_Cfunc_free
//go:linkname __cgofn__cgo_4dbd5e71a156_Cfunc_free _cgo_4dbd5e71a156_Cfunc_free
var __cgofn__cgo_4dbd5e71a156_Cfunc_free byte
var _cgo_4dbd5e71a156_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_4dbd5e71a156_Cfunc_free)

func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_4dbd5e71a156_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
