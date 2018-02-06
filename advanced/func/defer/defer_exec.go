package main

import(
	"fmt"
)

func double(x int)(result int){
	fmt.Println("before defer")
	//不到函数的最后一刻是不会执行的，能够改变外层函数返回给调用者的jie'guo
	defer func() {
		result *= 2
		fmt.Printf(" double(%d) =  %d \n", x , result)
	}()

	fmt.Println("after  defer")

	return x + x

}

func main() {
	fmt.Println("in main ")
	double(4)

	fmt.Println("after main ")

}