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
	imapClient, err := client.DialTLS("imap.gmail.com:993", nil)
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
		fmt.Println(inbox.Messages)
		// seqset := new(imap.SeqSet)
		// seqset.AddRange(inbox.Messages, inbox.Messages)
		// messages := make(chan *imap.Message, 1)
		// section := &imap.BodySectionName{}
		// if err := imapClient.Fetch(seqset, []imap.FetchItem{section.FetchItem()}, messages); err != nil {
		// 	log.Println(err)
		// }
		// msg:=<-messages
		// if msg==nil{
		// 	log.Println("Server didn't Messages")
		// }
		// r:=msg.GetBody(section)
		// if r==nil{
		// 	log.Println("Server didn't return message body")
		// }
		// ml,err:=mail.CreateReader(r)
		// if err!=nil{
		// 	log.Println(err)
		// }

	} else {
		fmt.Println("No new Message")
	}

}
