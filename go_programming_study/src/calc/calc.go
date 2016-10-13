package main

import "os"
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
	fmt.Println(" usage: calc command [arguments] ...")
	fmt.Println("\n the commands are:\n\tadd\taddition of two values.")
	fmt.Println("\n \tsqrt\tsquare root of a non-negative value.\n")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	fmt.Println("args[1]:", args[1], "len:", len(args))
	switch args[1] {

		case "add":
			if len(args) != 4 {
				fmt.Println("usage: calc add <int1> <int2>")
				return
			}
			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil  || err2 != nil {
				fmt.Println("usage: calc add <int1> <int2>")
				return
			}
			ret := simplemath.Add(v1, v2)
			fmt.Println("result:", ret)

		case "sqrt":
			if len(args) !=3 {
				fmt.Println("usage: calc sqrt <int1>")
				return
			}
			v, err := strconv.Atoi(args[2])
			if err != nil  {
				fmt.Println("usage: calc sqrt <int>")
				return
			}
			ret := simplemath.Sqrt(v)
			fmt.Println("result:", ret)

		default:
			Usage()
	}
}
