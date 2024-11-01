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
	FirstName     string `json:"FirstName"|nil`
	LastName      string `json:"LastName"|nil`
	EmailAddress  string `json:"EmailAddress"|nil`
	DecisionMaker bool   `json:"DecisionMaker"|nil`
	Country       string `json:"Country"|nil`
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
		err := json.NewDecoder(strings.NewReader(line)).Decode(&user)
		if err != nil {
			log.Println(err)
		}
		if user.EmailAddress != "" {
			fmt.Println(user)
		}
		slower += 1
		if slower%100 == 0 {

			time.Sleep(30 * time.Second)
		}
		if slower == 1000 {
			break
		}

	}
	println("found", count, "email")
}
