package main
import "fmt"
/*
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
(The data flows in the direction of the arrow.)

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.
 */
func sum(values [] int, resultChan chan int) { 
	sum := 0
	for _, value := range values { 
		sum += value
	}
	resultChan <- sum 
	// 将计算结果发送到channel中 
}
func main() {
	values := [] int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan // 接收结果
	fmt.Println("Result:", sum1, sum2, sum1 + sum2)
	fmt.Print("test")
}
