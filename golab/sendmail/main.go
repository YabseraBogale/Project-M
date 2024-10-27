package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
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
	to := []string{"yabserapython@gmail.com"}

	// SMTP server configuration
	smtpHost := "smtp.gmail.com" // Example: smtp.gmail.com
	smtpPort := "587"            // For SSL: 465, for STARTTLS: 587

	// Message
	subject := "Subject: Test HTML Email\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// Combine headers and body
	message := []byte(subject + mime + str)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send Email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully!")
	fmt.Println(str)
}
