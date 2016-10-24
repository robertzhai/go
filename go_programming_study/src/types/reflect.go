package main

import (
	"reflect"
	"fmt"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	//modify value
	p := reflect.ValueOf(&x)
	val := p.Elem()
	val.SetFloat(7.1)
	fmt.Println(x)
}
