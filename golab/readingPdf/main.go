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
	os.WriteFile("data.txt", file, 0777)

	fmt.Println("file written successfully.")
}
