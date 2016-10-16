package main

import (

	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error("A group of walrus emerges from the ocean")

}
