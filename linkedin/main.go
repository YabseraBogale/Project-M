package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Data struct {
	FirstName     string `json:"FirstName"`
	LastName      string `json:"LastName"`
	EmailAddress  string `json:"EmailAddress"`
	DecisionMaker string `json:"DecisionMaker"`
	Country       string `json:"Country"`
}

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(file))
	count := 0
	slower := 1
	var user Data
	for scanner.Scan() {
		line := scanner.Text()
		if slower%100 == 0 {
			err := json.NewDecoder(strings.NewReader(line)).Decode(&user)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(user)
			time.Sleep(30 * time.Second)
		}
		if slower == 2 {
			break
		}

	}
	println("found", count, "email")
}
