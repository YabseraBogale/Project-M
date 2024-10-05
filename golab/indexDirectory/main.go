package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	fileUrl := "http://maslovd.ru/public/mp3/Muse/Albums%20(CD)/2022%20-%20Will%20Of%20The%20People/06%20You%20Make%20Me%20Feel%20Like%20It's%20Halloween.mp3"

	// Download the file, params:
	// 1) name of file to save as
	// 2) URL to download FROM
	err := DownloadFile("saveas.mp3", fileUrl)
	if err != nil {
		fmt.Println("Error downloading file: ", err)
		return
	}

	fmt.Println("Downloaded: " + fileUrl)
}

// DownloadFile will download from a given url to a file. It will
// write as it downloads (useful for large files).
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
