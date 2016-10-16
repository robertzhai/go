package main

import (
	"fmt"
	"flag"
	"github.com/golang/glog"
)



/**
go run golog_debug.go -log_dir="./"

 */



// 避免没有引用fmt的编译错误
var _ = fmt.Println

func main() {
	//初始化命令行参数
	//初始化命令行参数
	flag.Parse()
	glog.Info("hello, glog")
	// 退出时调用，确保日志写入文件中
	glog.Flush()
}