package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/vova616/screenshot"
)

func main() {
	logfile, err := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(logfile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Println(err)
		}
		imgfile, err := os.Create("./file.png")
		if err != nil {
			log.Println(err)
		}
		defer imgfile.Close()
		err = png.Encode(imgfile, img)
		if err != nil {
			log.Println(err)
		}
		fmt.Scanf("Helloworld", w)
	})
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println(err)
	}
}
