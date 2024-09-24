package main

import (
	"fmt"
	"net/http"
)

func HelloJavascript(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloJavascript)
	http.ListenAndServe(":8000", nil)
}
