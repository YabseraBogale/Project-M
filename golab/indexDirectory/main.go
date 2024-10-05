package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	url := "http://maslovd.ru/public/mp3/Muse/Albums%20(CD)/?C=N;O=D"
	webscraper := colly.NewCollector()

	webscraper.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL)
	})

	err := webscraper.Visit(url)
	if err != nil {
		log.Println(err)
	}
}
