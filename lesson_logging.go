package main

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func _main() {
	_, err := os.Open("fsadaa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Printf("%T %v", "test", "test")
}
