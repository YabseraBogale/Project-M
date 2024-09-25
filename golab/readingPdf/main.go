package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	file, err := os.ReadFile("sample.pdf")
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "stream" {
			count = 1
		} else if count == 1 && line != "endstream" {
			fmt.Println(utf8.DecodeLastRuneInString(line))
		} else if line == "endstream" {
			count = 0
		}

	}

}
