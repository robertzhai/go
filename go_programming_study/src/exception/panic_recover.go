package main

import "fmt"

func main(){

	fmt.Println("the start of main")
	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
		fmt.Println("the end of func")
	}()
	f()
	fmt.Println("the end of main")
}

func f(){
	fmt.Println("the start of f ")
	panic(55)
	fmt.Println("the end of f")
}
