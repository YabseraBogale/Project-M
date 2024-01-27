package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.ReadDir(".")
	if err != nil {
		println("Failed in the beginnig")
	}
	for _, i := range file {
		err = filepath.Walk(i.Name(), func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				println("Failed at handling path")
			}
			if info.IsDir() && info.Name() == "Skip" {
				println("skipping a dir without errors: %+v \n", info.Name())
				return filepath.SkipDir
			}
			println("path at " + path)
			return nil
		})
	}
}
