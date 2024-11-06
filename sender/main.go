package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "gold.db")
	if err != nil {
		log.Println("err with open", err)
	}
	defer db.Close()
	state, err := db.Query(`select Country from UserData`)
	if err != nil {
		log.Println("err with query", err)
	}
	Indonesia := 0
	Germany := 0
	Saudi := 0
	Angola := 0
	USA := 0
	Unknown := 0
	count := 0

	for state.Next() {
		var Country string
		count += 1
		err := state.Scan(&Country)
		if err != nil {
			log.Println("err with scan", err)
		}
		if Country == "Indonesia" {
			Indonesia += 1
		} else if Country == "Germany" {
			Germany += 1
		} else if Country == "Saudi" {
			Saudi += 1
		} else if Country == "Angola" {
			Angola += 1
		} else if Country == "USA" {
			USA += 1
		} else {
			Unknown += 1
		}

	}

	indonesia := float32((Indonesia * 100)) / float32(count)
	germany := float32((Germany * 100)) / float32(count)
	saudi := float32((Saudi * 100)) / float32(count)
	angola := float32((Angola * 100)) / float32(count)
	usa := float32((USA * 100)) / float32(count)
	unknown := float32((Unknown * 100)) / float32(count)
	ExpectedPeopleToSeeMin := count * 17 / 100
	ExpectedPeopleToSeeMax := count * 25 / 100
	ExpectedPeopleToSeeCallMin := ExpectedPeopleToSeeMin * 17 / 100
	ExpectedPeopleToSeeCallMax := ExpectedPeopleToSeeMin * 25 / 100
	DataSent := float32(10*count) / float32(1000000)
	fmt.Println("All Assumations being all the email works")
	fmt.Println("Indonesia:", Indonesia, "Percentage:", indonesia)
	fmt.Println("Germany:", Germany, "Percentage:", germany)
	fmt.Println("Saudi:", Saudi, "Percentage:", saudi)
	fmt.Println("Angola:", Angola, "Percentage:", angola)
	fmt.Println("USA:", USA, "Percentage:", usa)
	fmt.Println("Unknown:", Unknown, "Percentage:", unknown)
	fmt.Println("Expected 17% number of people to see the email", ExpectedPeopleToSeeMin)
	fmt.Println("Expected 25% number of people to see the email", ExpectedPeopleToSeeMax)
	fmt.Println("Expected 17% number of people to call", ExpectedPeopleToSeeCallMin)
	fmt.Println("Expected 25% number of people to call", ExpectedPeopleToSeeCallMax)
	fmt.Println("Size of data sent over email for", count, "is", DataSent, "gb")

}
