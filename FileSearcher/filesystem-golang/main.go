package main

import (
	"os"
)

func main() {

	file, err := os.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		ff, err := os.ReadDir(i.Name())
		if err != nil {
		}
		for _, i := range ff {
			println(i.IsDir())

		}
	}

}
