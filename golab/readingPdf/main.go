package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("sample.pdf")
	if err != nil {
		log.Println(err)
	}
	for _, i := range file {
		fmt.Println(i)
	}
}
