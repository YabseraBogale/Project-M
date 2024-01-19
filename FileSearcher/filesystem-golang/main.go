package main

import (
	"os"
	"strings"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {

	}
	var listdir []string = strings.Split(dir, "/")
	for i := 0; i < len(listdir); i++ {
		println(listdir[i])
	}

}
