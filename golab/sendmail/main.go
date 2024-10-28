package main

import (
	"bytes"
	"gopkg.in/mail.v2"
	"html/template"
	"log"
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
	from := "yabsera@frielvh.com"
	password := "YSFYy&h{{_Fn"
	to := "yabserapython@gmail.com"

	// SMTP server configuration
	smtpHost := "mail.frielvh.com" // Example: smtp.gmail.com
	smtpPort := 587                // For SSL: 465, for STARTTLS: 587
	m := mail.NewMessage()

	// Set the sender and recipient addresses
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Test HTML Email")

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
