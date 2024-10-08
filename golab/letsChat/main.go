package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, err := template.ParseFiles("public/index.html")
		if err != nil {
			log.Println(err)
		}
		err = temp.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	})
	http.ListenAndServe(":8080", nil)

}
