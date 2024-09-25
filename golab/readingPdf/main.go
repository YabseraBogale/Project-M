package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("sample.pdf")
	if err != nil {
		log.Println(err)
	}
	os.WriteFile("data.txt", file, 0777)
	//fmt.Println(string(file[8:9]) == "\n")
	result := bytes.Split(file, []byte("stream"))
	for _, i := range result {
		fmt.Println(string(i))
	}
}
