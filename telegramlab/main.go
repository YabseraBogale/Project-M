package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zelenin/go-tdlib/client"
)

func main() {
	// Set up the configuration for TDLib client
	apiID := os.Getenv("API_ID")             // Replace with your API ID from my.telegram.org
	apiHash := os.Getenv("API_HASH")         // Replace with your API Hash from my.telegram.org
	phoneNumber := os.Getenv("PHONE_NUMBER") // Your phone number

	tdlibClient := client.NewClient(&client.Config{
		APIID:              apiID,
		APIHash:            apiHash,
		SystemLanguageCode: "en",
		DeviceModel:        "Desktop",
		SystemVersion:      "1.0",
		ApplicationVersion: "1.0",
	})

	// Authorization process
	tdlibClient.AuthSetPhoneNumber(phoneNumber)

	for {
		currentState, err := tdlibClient.AuthGetState()
		if err != nil {
			log.Fatalf("Failed to get auth state: %v", err)
		}

		// If we need a login code
		if currentState.GetAuthorizationStateEnum() == client.AuthorizationStateWaitCodeType {
			var code string
			fmt.Println("Enter the login code you received via Telegram: ")
			fmt.Scanln(&code)

			// Send the code to Telegram
			err = tdlibClient.AuthSetCode(code)
			if err != nil {
				log.Fatalf("Failed to set login code: %v", err)
			}
		}

		if currentState.GetAuthorizationStateEnum() == client.AuthorizationStateReadyType {
			break
		}

		time.Sleep(1 * time.Second)
	}

	// Now you're logged in

	// Send a message to yourself
	chat, err := tdlibClient.SearchPublicChat("me")
	if err != nil {
		log.Fatalf("Failed to find chat: %v", err)
	}

	messageText := &client.InputMessageText{
		Text: &client.FormattedText{
			Text: "Hello from Go!",
		},
	}

	_, err = tdlibClient.SendMessage(&client.SendMessageRequest{
		ChatID:              chat.ID,
		InputMessageContent: messageText,
	})
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Println("Message sent successfully!")
}

