package main

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/mail.v2"
)

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

func main() {

	str, err := GetHtml()
	if err != nil {
		log.Fatalln(err)
	}
	from := "cheretaaddis@gmail.com"
	password := "zgzd xtlt emlc tzfb"
	to := "yabserapython@gmail.com"

	// SMTP server configuration
	smtpHost := "smtp.gmail.com" // Example: smtp.gmail.com
	smtpPort := 587              // For SSL: 465, for STARTTLS: 587
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
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
