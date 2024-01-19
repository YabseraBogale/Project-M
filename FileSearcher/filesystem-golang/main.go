package main

import "os"

func main() {
	dir, err := os.ReadDir(".")
	if err != nil {

	}
	for _, name := range dir {
		print(name)
	}
}
