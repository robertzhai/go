package main

import "fmt"

func main()  {
	fmt.Print("start ")
	panic("panic")
	fmt.Print("end")
}
