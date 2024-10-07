package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

	var wg sync.WaitGroup
	listOfLinkd := make(map[string]string)
	for _, i := range data[1:] {
		wg.Add(1)
		go func(src string) {
			defer wg.Done()
			muses := colly.NewCollector()
			muses.OnHTML("td", func(h *colly.HTMLElement) {
				if strings.Count(h.ChildAttr("a", "href"), "mp3") >= 1 && h.ChildAttr("a", "href") != "/public/mp3/Muse/Albums%20(CD)/" {
					filename := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(h.ChildAttr("a", "href"), "%20", " "), "%5b", ""), "%5d", "")
					listOfLinkd[filename] = src + h.ChildAttr("a", "href")
				}
			})
			err := muses.Visit(src)
			if err != nil {
				log.Println(err)
			}

		}(i)

	}
	wg.Wait()

	for i := range listOfLinkd {
		file, err := os.Create("Downloads/" + i)
		if err != nil {
			log.Println(err)
			break
		}
		defer file.Close()
		client := http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				req.URL.Opaque = req.URL.Path
				return nil
			},
		}
		res, err := client.Get(listOfLinkd[i])
		if err != nil {
			log.Println(err)
			break
		}
		defer res.Body.Close()
		size, err := io.Copy(file, res.Body)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(i, size/10000)
	}
	fmt.Println("All Operation Completed")
}
