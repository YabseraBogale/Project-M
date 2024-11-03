package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	userdata := map[string][]string{}
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Println("err with open", err)
	}
	defer db.Close()
	state, err := db.Query(`select distinct Firstname,Lastname,Email,Country,HQCompanyName from useremail`)
	if err != nil {
		log.Println("err with query", err)
	}
	for state.Next() {
		var Firstname string
		var Lastname string
		var EmailAddres string
		var Country string
		var HQCompanyName string
		err := state.Scan(&Firstname, &Lastname, &EmailAddres, &Country, &HQCompanyName)
		if err != nil {
			log.Println("err with scan", err)
		}
		userdata[EmailAddres] = []string{Firstname, Lastname, Country, HQCompanyName}
	}
	Indonesia := 0
	Germany := 0
	Saudi := 0
	Unknown := 0
	count := 0
	for _, i := range userdata {
		count += 1
		if i[2] == "Saudi" {
			Saudi += 1
		} else if i[2] == "Germany" {
			Germany += 1
		} else if i[2] == "Indonesia" {
			Indonesia += 1
		} else {
			Unknown += 1
		}
	}
	indonesia := float32((Indonesia * 100)) / float32(count)
	germany := float32((Germany * 100)) / float32(count)
	saudi := float32((Saudi * 100)) / float32(count)
	unknown := float32((Unknown * 100)) / float32(count)
	ExpectedPeopleToSeeMin := count * 17 / 100
	ExpectedPeopleToSeeMax := count * 25 / 100
	ExpectedPeopleToSeeCallMin := ExpectedPeopleToSeeMin * 17 / 100
	ExpectedPeopleToSeeCallMax := ExpectedPeopleToSeeMin * 25 / 100
	fmt.Println("All Assumations being all the email works")
	fmt.Println("Indonesia:", Indonesia, "Percentage:", indonesia)
	fmt.Println("Germany:", Germany, "Percentage:", germany)
	fmt.Println("Saudi:", Saudi, "Percentage:", saudi)
	fmt.Println("Unknown:", Unknown, "Percentage:", unknown)
	fmt.Println("Expected 17% number of people to see the email", ExpectedPeopleToSeeMin)
	fmt.Println("Expected 25% number of people to see the email", ExpectedPeopleToSeeMax)
	fmt.Println("Expected 17% number of people to call", ExpectedPeopleToSeeCallMin)
	fmt.Println("Expected 25% number of people to call", ExpectedPeopleToSeeCallMax)

}
