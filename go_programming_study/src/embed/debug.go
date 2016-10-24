package main

import "fmt"

type debug struct {

}

func (d debug) print()  {
	fmt.Println("debug")
}

type file struct {
	debug
}

func main() {
	var f file
	f.print()
}
