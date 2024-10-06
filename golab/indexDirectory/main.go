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

	for _, i := range data[1:] {
		wg.Add(1)
		go func(src string) {
			defer wg.Done()
			muses := colly.NewCollector()
			muses.OnHTML("td", func(h *colly.HTMLElement) {
				if strings.Count(h.ChildAttr("a", "href"), "mp3") >= 1 && h.ChildAttr("a", "href") != "/public/mp3/Muse/Albums%20(CD)/" {
					filename := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(h.ChildAttr("a", "href"), "%20", " "), "%5b", ""), "%5d", "")
					file, err := os.Create("Downloads/" + filename)
					if err != nil {
						log.Println(err, "for", filename)
					}
					defer file.Close()
					client := http.Client{
						CheckRedirect: func(req *http.Request, via []*http.Request) error {
							req.URL.Opaque = req.URL.Path
							return nil
						},
					}
					res, err := client.Get(src + h.ChildAttr("a", "href"))
					if err != nil {
						log.Println(err, "for", filename)
					}
					defer res.Body.Close()
					size, err := io.Copy(file, res.Body)
					if err != nil {
						log.Println(err, "for", filename)
						log.Println(src + h.ChildAttr("a", "href"))
						err := os.Remove("Downloads/" + filename)
						if err != nil {
							log.Println(err)
						}
					}
					fmt.Println("Downloaded a file", filename, " with size", size/100000)
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
