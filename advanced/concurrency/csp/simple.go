package main;

import(
	"fmt"
)

var a string
var done bool

func setup() {
	a = "hello world\n"
	done = true;
}

func main() {

	go setup()
	for !done {
      	fmt.Println("in main \n")
	}
	print(a)
	
}