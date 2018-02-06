package main

import(
	"fmt"
	"sync"
)


func main() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello mutex sync in another go routine \n")
	}()
	mu.Unlock()

  	fmt.Println("in main \n")
	
}