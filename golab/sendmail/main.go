package main

import (
	"bytes"
	"database/sql"
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
	db, err := sql.Open("sqlite3", "gold.db")
	if err != nil {
		log.Println("err with open", err)
	}
	defer db.Close()

	EmailSender("yabserapython@gmail.com")
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
