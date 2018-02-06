package main

import(
	"fmt"
)

func main() {

	fmt.Print("main start ");
	 naturals := make(chan int)
	 squares := make(chan int)
	 //counter
	 go func() {
		 for x := 0; ; x++ {
			 naturals <- x
		 }
	 }()

	 go func() {
		for {
			x := <- naturals
			squares <- x * x
		}
	 }()

	 // print 
	 for {
		 fmt.Println(<- squares)
	 }




}