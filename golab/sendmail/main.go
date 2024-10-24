package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

func GetHtml() (string, error) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		return "", err
	}
	f := []string{}
	buf := new(bytes.Buffer)
	if err = temp.Execute(buf, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}

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
