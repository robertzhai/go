package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "os"
import "strconv"
import "sorter/algorithms/bubblesort"
import (
	"sorter/algorithms/qsort"
	"time"
)

var infile *string = flag.String("i", "infile", "file contains data to sort")
var outfile *string = flag.String("o", "outfile", " file to save sorted data")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func readValues(infile string)(values []int, err error)  {

	file, err := os.Open(infile)
	if err != nil {
		fmt.Println(" failed to open file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("a too long line")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return

}

func writeValues(values [] int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("failed to create ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil

}

func main() {

	flag.Parse()

	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=",
			*algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {
		t1 := time.Now()
		switch *algorithm {

		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("algorithm ", *algorithm, " unknown")

		}
		t2 := time.Now()
		fmt.Println("time cost: " , t2.Sub(t1), " to complete")
		writeValues(values, *outfile)
		fmt.Println("read values:", values)
	} else {

		fmt.Println(err)
	}


}
