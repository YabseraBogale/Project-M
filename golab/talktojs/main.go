package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm")
				return
			}
			fmt.Fprintf(w, "post from website\n")
			username := r.FormValue("username")
			fmt.Fprintf(w, "Name = %s\n", username)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")

		}
	}
}

func main() {
	f := http.FileServer(http.Dir("./static"))
	http.Handle("/", f)
	http.HandleFunc("/hello", hello)
	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
