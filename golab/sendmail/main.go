package main

import (
	"log"
	"net/smtp"
	"os"
)

func main() {
	smtphost := "smtp.gmail.com"
	smtpport := "587"
	email := "cheretaaddis@gmail.com"
	to := []string{"yabserapython@gmail.com"}
	app_id := os.Getenv("app_id")
	msg := "subject: This is to test golang\n\nthis the body to the message"

	message := []byte(msg)
	auth := smtp.PlainAuth("", email, app_id, smtphost)
	err := smtp.SendMail(smtphost+":"+smtpport, auth, email, to, message)
	if err != nil {
		log.Fatalln(err)
	}
}
