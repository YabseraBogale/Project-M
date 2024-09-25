package main

import (
	"bufio"
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
	var data []byte
	count := 0
	for i := 0; i < len(file); i++ {
		line, _, err := scanner.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		if string(line) == "stream" {
			count = 1
		} else if count == 1 && string(line) != "endstream" {
			data = append(data, line...)
		} else if string(line) == "endstream" {
			count = 0
		}

	}
	os.WriteFile("data.txt", data, 0777)

}
