package main

import (
	"bufio"
	"compress/zlib"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("sample.pdf")
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	count := 0
	data := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "stream" {
			count = 1
		} else if count == 1 && line != "endstream" {
			r, err := zlib.NewReader(strings.NewReader(line))
			if err != nil {
				println("here")
				log.Println(err)
			}
			b := make([]byte, len(line))
			for {
				n, err := r.Read(b)
				data = append(data, b[:n])
				if err != nil {
					return
				}

			}

		} else if line == "endstream" {
			count = 0
		}

	}
	os.WriteFile("data.txt", data[0], 0777)
}
