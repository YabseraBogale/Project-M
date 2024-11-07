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

}
