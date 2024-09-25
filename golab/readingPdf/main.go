package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("sample.pdf")
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewReader(strings.NewReader(string(file)))
	for i := 0; i < len(file); i++ {
		line, _, err := scanner.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(string(line))

	}
}
