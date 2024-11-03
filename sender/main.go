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
	state, err := db.Query(`select distinct Firstname,Lastname,Email,Country from useremail`)
	if err != nil {
		log.Println("err with query", err)
	}
	for state.Next() {
		var Firstname string
		var Lastname string
		var EmailAddres string
		var Country string
		err := state.Scan(&Firstname, &Lastname, &EmailAddres, &Country)
		if err != nil {
			log.Println("err with scan", err)
		}
		userdata[EmailAddres] = []string{Firstname, Lastname, Country}
	}
	fmt.Println(userdata)
}
