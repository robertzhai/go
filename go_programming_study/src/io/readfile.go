package main

import (
	"io"
	"os"
	"fmt"
)

func init() {
	fmt.Println("first init")
}

func init() {
	fmt.Println("second init")
}

func main() {
	content, err := Contents("./data.txt")
	if err != nil {
		fmt.Println("read err: %v", err)
		return
	}
	fmt.Print("file contents:", content)
}

// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()  // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:100])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err  // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}