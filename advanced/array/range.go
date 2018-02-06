package main 

import "fmt"

func main() {
	data := []int {1,3,5}

	for i,value := range data {

		fmt.Println(i)
		fmt.Println(value)

	}
}