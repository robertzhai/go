package main

import "fmt"

type Rectangle struct{
	height, weight float64
}

func main()  {

	var r0 Rectangle
	r0.height = 5
	r0.weight = 3
	fmt.Println(r0)
	r1 := Rectangle{1,2}
	fmt.Println(r1)
	r2 :=Rectangle{height:5, weight:3}
	fmt.Println(r2)
	r3 := new(Rectangle)
	r3.height = 5
	r3.weight = 3
	fmt.Println(r3)


}
