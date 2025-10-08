package chopstick

import (
	"log"
	"os"
)

var debug *log.Logger

func LogInit() {
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal("Error opening log file", err)
		os.Exit(1)
	}

	debug = log.New(f, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
