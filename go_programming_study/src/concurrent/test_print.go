package main

import (
	"fmt"
	"time"
)

func print(s string) {
	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go print("world")
	print("hello")
}