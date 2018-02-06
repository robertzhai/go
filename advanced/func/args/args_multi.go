package main 

import(
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {

	arr := []int{1,3,5,7,9}
	fmt.Print(sum(1,3,5), arr)
	fmt.Print(sum(arr...))


}