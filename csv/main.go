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
		fmt.Println("Reading db", err)
	}

	statment, err := db.Query(`Select * from UserData Where Country="USA";`)
	if err != nil {
		log.Println("Prepare statemnt", err)
	}

}
