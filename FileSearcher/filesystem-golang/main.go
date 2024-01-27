package main

import "path/filepath"

func main() {
	err := filepath.Walk(".", func() {

	})
}
