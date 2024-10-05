package main

import (
	"fmt"
	"log"
	"sync"

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

	for _, i := range data[1:] {
		var wait sync.WaitGroup
		go func(src string) {
			wait.Wait()
			muses := colly.NewCollector()
			muses.OnHTML("td", func(h *colly.HTMLElement) {
				fmt.Println(h.ChildAttr("a", "href"))
			})
			err := muses.Visit(src)
			if err != nil {
				log.Println(err)
			}
		}(i)
		wait.Done()

	}
}
