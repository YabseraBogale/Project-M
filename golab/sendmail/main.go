package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func GetHtml() (string, error) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		return "", err
	}
	f := []string{"Hello World"}
	buf := new(bytes.Buffer)
	if err = temp.Execute(buf, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {
	str, err := GetHtml()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(str)
}
