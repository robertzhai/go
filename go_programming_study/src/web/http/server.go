package main

import (
	"io"
	"net/http"
	"log"
	_ "net/http/pprof"
)

// hello go, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, go!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}