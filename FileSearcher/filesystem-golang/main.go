package main

import (
	"os"
)

func main() {

	file, err := os.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		ff, err := os.ReadDir(i.Name() + "/.")
		if err != nil {
		}
		for _, j := range ff {
			println(j.Name())
		}
	}

}
