package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	FirstName     string `json:"FirstName"|nil`
	LastName      string `json:"LastName"|nil`
	EmailAddress  string `json:"Email"|nil`
	DecisionMaker bool   `json:"DecisionMaker"|nil`
	Country       string `json:"Country"|nil`
	HQCompanyName string `json:"HQCompanyName"|nil`
}

const filename string = "database.db"

func main() {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Println(err)
	}

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
			statment := `Insert into UserEmail(FirstName,LastName,Country,Email,HQCompanyName) values(?,?,?,?,?)`
			state, err := db.Prepare(statment)
			if err != nil {
				log.Println(err)
			}
			_, err = state.Exec(user.FirstName, user.LastName, user.Country, user.EmailAddress, user.HQCompanyName)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Sucessfully Inserted line at .... ", slower)
		}
		slower += 1
		if slower%100 == 0 {

			time.Sleep(2 * time.Second)
		}

	}

}
