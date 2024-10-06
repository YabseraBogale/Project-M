package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	println("The Number Of Data Points", len(data[1:]))
	var wg sync.WaitGroup

	for _, i := range data[1:] {
		wg.Add(1)
		go func(src string) {
			defer wg.Done()
			muses := colly.NewCollector()
			muses.OnHTML("td", func(h *colly.HTMLElement) {
				if strings.Count(h.ChildAttr("a", "href"), "mp3") >= 1 && h.ChildAttr("a", "href") != "/public/mp3/Muse/Albums%20(CD)/" {
					filename := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(h.ChildAttr("a", "href"), "%20", " "), "%5b", ""), "%5d", "")
					_, err := os.Create("Downloads/" + filename)
					if err != nil {
						log.Println(err)
					}
				}
			})
			err := muses.Visit(src)
			if err != nil {
				log.Println(err)
			}

		}(i)

	}
	wg.Wait()
	fmt.Println("All Operation are Completed")
}
