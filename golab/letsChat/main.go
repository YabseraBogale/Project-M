package main

import (
	"net/http"
)

func main() {
	file := http.FileServer(http.Dir("public"))
	http.Handle("/", file)
	http.ListenAndServe(":8080", nil)

}
