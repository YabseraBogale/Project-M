package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/vova616/screenshot"
)

func main() {
	logfile, err := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.SetOutput(logfile)
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	upload := fmt.Sprintf("%s/.uploads", home)
	err = os.Mkdir(upload, 0666)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	for {
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Println(err)
			panic(err)
		}

		filename := fmt.Sprintf("%s/.%s.png", upload, time.Now())
		imgfile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		err = png.Encode(imgfile, img)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		time.Sleep(3 * time.Second)
		imgfile.Close()
	}
}
