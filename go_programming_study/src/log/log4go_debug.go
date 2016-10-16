package main

import (
	"github.com/alecthomas/log4go"
)

func main() {

	logger3 := log4go.Logger{}

	logger3.AddFilter("stdout", log4go.DEBUG, log4go.NewConsoleLogWriter())
	logger3.AddFilter("file", log4go.DEBUG, log4go.NewFileLogWriter("file.log", false))

	logger3.Debug("debug info")
	logger3.Error("debug info")
	defer logger3.Close()

}
