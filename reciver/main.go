package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emersion/go-imap/client"
)

const username string = "cheretaaddis@gmail.com"

var password string = os.Getenv("password")

func main() {
	file, err := os.OpenFile("reciver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	imapClient, err := client.DialTLS("imap.google.com", nil)
	if err != nil {
		log.Println(err)
	}
	defer imapClient.Logout()

	if err := imapClient.Login(username, password); err != nil {
		log.Println(err)
	}
	fmt.Println("Logged in successfully")
	inbox, err := imapClient.Select("INBOX", false)
	if err != nil {
		log.Println(err)
	}
	if inbox.Messages > 0 {

	} else {
		fmt.Println("No new Message")
	}

}
