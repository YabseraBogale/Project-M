package main

import (
	"os"
)

func main() {

	file, err := os.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		print(i.Name())
	}

}
