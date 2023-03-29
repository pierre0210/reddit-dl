package util

import (
	"log"
	"os"
)

func ErrHandler(err error, fatal bool) {
	if err == nil {
		return
	} else if fatal {
		log.Fatalf("%s\n", err.Error())
	} else {
		log.Printf("%s\n", err.Error())
		os.Exit(0)
	}
}
