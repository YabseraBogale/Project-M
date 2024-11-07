package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	go_mail "github.com/emersion/go-message/mail"
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

	imapClient, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer imapClient.Logout()

	if err := imapClient.Login(from, password); err != nil {
		log.Println("Login failed:", err)
	}

	// If login is successful
	fmt.Println("Logged in successfully!")

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

			mailbox, err := imapClient.Select("INBOX", false)
			if err != nil {
				log.Println(err)
			}
			if mailbox.Messages > 0 {
				seqset := new(imap.SeqSet)
				seqset.AddRange(mailbox.Messages, mailbox.Messages)
				messages := make(chan *imap.Message, 1)
				section := &imap.BodySectionName{}
				err = imapClient.Fetch(seqset, []imap.FetchItem{section.FetchItem()}, messages)
				if err != nil {
					log.Println(err)
				}
				msg := <-messages
				if msg == nil {
					log.Println("Server didn't return message")
				}
				r := msg.GetBody(section)
				if r == nil {
					log.Println("Server didn't return message body")
				}
				mr, err := go_mail.CreateReader(r)
				if err != nil {
					log.Fatal(err)
				}
				header := mr.Header
				from, err := header.AddressList("From")
				if err == nil {
					fmt.Println("From:", from)
				}
				subject, err := header.Subject()
				if err == nil {
					fmt.Println("Subject:", subject)
				}
			}
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
