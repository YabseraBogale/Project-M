package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("SC.csv")
	if err != nil {
		log.Println("File Reading", err)
	}

	read, err := csv.NewReader(bytes.NewReader(file)).ReadAll()
	if err != nil {
		log.Println("csv parser", err)
	}
	for _, i := range read {
		fmt.Println(i)
	}

}
