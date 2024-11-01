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
	HQCompanyName string `json:"HQCompanyName"|nil`
}

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Println(err)
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(file))
	slower := 1
	var user Data
	for scanner.Scan() {
		line := scanner.Text()
		err := json.NewDecoder(strings.NewReader(line)).Decode(&user)
		if err != nil {
			log.Println(err)
		}
		if user.EmailAddress != "" && user.DecisionMaker == true {
			fmt.Println("At line", slower, user)
		}
		slower += 1
		if slower%100 == 0 {

			time.Sleep(30 * time.Second)
		}
		if slower == 10000 {
			break
		}

	}

}
