package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/mail.v2"
)

const from string = "cheretaaddis@gmail.com"

const smtpHost = "smtp.gmail.com" // Example: smtp.gmail.com
const smtpPort = 587

var password string = os.Getenv("password")

func main() {

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)



	// If login is successful

	db, err := sql.Open("sqlite3", "gold.db")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	// 	select Email,Sent from userdata WHERE Country='Germany' and Sent='not_sent' LIMIT 50; -- sent
	// 	select Email,Sent from userdata WHERE Country='Saudi' and Sent='not_sent' LIMIT 50;
	// 	select Email,Sent from userdata WHERE Country='USA' and Sent='not_sent' LIMIT 50;
	// 	select Email,Sent from userdata WHERE Country='Indonesia' and Sent='not_sent' LIMIT 50;
	// 	select Email,Sent from userdata WHERE Country='Angola' and Sent='not_sent' LIMIT 50;
	row, err := db.Query(`Select Email From userdata WHERE Country='Angola' and Sent='not_sent' LIMIT 4;`)
	if err != nil {
		log.Println(err)
	}
	defer row.Close()

	for row.Next() {

		var Email string

		row.Scan(&Email)

		sent, err := EmailSender(Email)
		if err != nil {
			log.Println(err)
		} else if sent {
			fmt.Println("Sent to...", Email)



			}
		}

}

func EmailSender(to string) (bool, error) {
	str, err := GetHtml()
	if err != nil {
		return false, err
	}
	// SMTP server configuration
	// For SSL: 465, for STARTTLS: 587
	m := mail.NewMessage()
	// Set the sender and recipient addresses
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Software Developer Application- Yabsera Bogale")

	// Set the email body as HTML
	m.SetBody("text/html", str)

	// SMTP server configuration
	d := mail.NewDialer(smtpHost, smtpPort, from, password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}

func GetHtml() (string, error) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		return "", err
	}
	f := []string{"Hello World"}
	buf := new(bytes.Buffer)
	if err = temp.Execute(buf, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}
