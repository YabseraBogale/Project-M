package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vova616/screenshot"
)

func main() {
	logfile, err := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(logfile)
	for {
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Println(err)
		}
		filename := fmt.Sprintf("%s.png", time.Now())
		imgfile, err := os.Create("")
	}
}
