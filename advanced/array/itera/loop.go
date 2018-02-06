package main

import(
	"fmt"
)

func main() {
	medals := []string{"gold","sliver", "bronze"}

	for i:= len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}


