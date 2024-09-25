package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
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
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		log.Println(err)
	}
	defer r.Close()
	enflated, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enflated))

}
