package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	url := "http://maslovd.ru/public/mp3/Muse/Albums%20(CD)/?C=N;O=D"
	album := "http://maslovd.ru/public/mp3/Muse/Albums%20(CD)"
	webscraper := colly.NewCollector()
	data := []string{}
	webscraper.OnHTML("td", func(h *colly.HTMLElement) {
		a := h.ChildAttr("a", "href")
		if a != "" {
			data = append(data, album+"/"+a)
		}
	})
	err := webscraper.Visit(url)
	if err != nil {
		log.Println(err)
	}
	ch := make(chan string)
	println("The Number Of Data Points", len(data[1:]))
	for _, i := range data[1:] {
		go func(src string) {
			muses := colly.NewCollector()
			muses.OnHTML("td", func(h *colly.HTMLElement) {
				ch <- h.ChildAttr("a", "href")
			})
			err := muses.Visit(src)
			if err != nil {
				log.Println(err)
			}
			close(ch)
		}(i)
		for value := range ch {
			if value != "" {
				fmt.Println("Recived", value)
			}
		}
	}

}
