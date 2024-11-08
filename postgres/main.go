package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

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
	postgres, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Println(err)
	}
	defer postgres.Close()

	row, err := sqlite.Query(`Select * from UserData;`)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	if err != nil {
		log.Println(err)
	}
	slower := 0

	for row.Next() {
		tx := postgres.MustBegin()
		slower += 1
		var Email string
		var Firstname string
		var Lastname string
		var Country string
		var HQCompanyName string
		var Sent string
		err := row.Scan(&Email, &Firstname, &Lastname, &Country, &HQCompanyName, &Sent)

		if err != nil {
			log.Println(err)
		}

		tx.MustExec("Insert into public.UserData(Email, Firstname, Lastname, Country, HQCompanyName, Sent) values($1, $2, $3, $4, $5, $6)", Email, Firstname, Lastname, Country, HQCompanyName, Sent)

		err = tx.Commit()
		if err != nil {
			log.Println(err)
		}
		if slower%1000 == 0 {
			fmt.Println("At line row ... ", slower)
			time.Sleep(2 * time.Second)
		}
	}

}
