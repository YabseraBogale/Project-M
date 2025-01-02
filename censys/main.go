package main

import (
	"log"
	"context"

	"github.com/ns3777k/go-shodan/v4/shodan" // go modules required
)

func main() {
	client := shodan.NewEnvClient(nil)
	dns, err := client.GetDNSResolve(context.Background(), []string{"google.com", "ya.ru"})

	if err != nil {
		log.Panic(err)
	} else {
		log.Println(dns["google.com"])
	}
}
