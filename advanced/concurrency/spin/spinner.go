package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main")
	go spinner(100 * time.Millisecond)
	fmt.Println("after spinner")

	const n  = 45
	fibN := fib(n) 
	fmt.Printf("fib(%d) = %d \n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1 ) + fib(x - 2)
}