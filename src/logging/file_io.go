package logging

import (
	"io"
	"log"
	"os"
)

func loggingSetting(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

/*
 * output log
 */
func Logging() {
	loggingSetting("test.log")
	_, err := os.Open("not_exists_tmp_file.log")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Printf("%T %v", "test", "test")
}
