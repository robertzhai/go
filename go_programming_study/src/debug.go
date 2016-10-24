package main

import "fmt"

func main()  {
	var array [10]int
	slice := array[2:4]
	fmt.Println(len(slice), "---", cap(slice))
	
}
