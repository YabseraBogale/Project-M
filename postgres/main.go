package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "Gold"
)

var password string = os.Getenv("dbpassword")

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	sqlite, err := sql.Open("sqlite3", "gold.db")
	if err != nil {
		log.Println(err)
	}
	defer sqlite.Close()
	postgres, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Println(err)
	}
	defer postgres.Close()

	row, err := sqlite.Query(`Select * from UserData limit 10;`)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	for row.Next() {
		var Email string
		var Firstname string
		var Lastname string
		var Country string
		var HQCompanyName string
		var Sent string
		err := row.Scan(&Email, &Firstname, &Lastname, &Country, HQCompanyName, Sent)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(Email, Firstname, Lastname, Country, HQCompanyName, Sent)
	}
}
