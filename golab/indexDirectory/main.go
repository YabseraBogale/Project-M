package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	for i := range listOfLinkd {
		wg.Add(1)
		go func(filename string, src string) {
			defer wg.Done()
			file, err := os.Create("Downloads/" + filename)
			if err != nil {
				log.Println(err)
			}
			defer file.Close()
			client := http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					req.URL.Opaque = req.URL.Path
					return nil
				},
			}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, src, nil)
			if err != nil {
				log.Println(err)
			}
			res, err := client.Do(req)
			if err != nil {
				log.Println(err)
				return
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				log.Printf("failed to download file, status code: %d", res.StatusCode)
			}
			size, err := io.Copy(file, res.Body)
			if err != nil {
				log.Printf("failed to download file, status code: %d", res.StatusCode)
			} else {
				mb := size / 10000
				fmt.Println(filename, mb, "mb")
			}

		}(i, listOfLinkd[i])
	}
	wg.Wait()
	fmt.Println("All Operation Completed")
}
