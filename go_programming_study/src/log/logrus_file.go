package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"fmt"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	wi, err:= os.OpenFile("logrus3.log", os.O_RDWR | os.O_CREATE| os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("err:", err)
		return
	}
	defer wi.Close()

	log.SetOutput(wi)

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error("A group of walrus emerges from the ocean")

}