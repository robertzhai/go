package main

import(
	"fmt"
)


func main() {

	done := make(chan int)

	go func() {
		fmt.Println("hello channel sync in another go routine \n")
		done <- 1
	}()

	<- done

  	fmt.Println("in main \n")
	
}