package main

import (
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		if i.IsDir() == true {
			ioutil.ReadAll(i.Name())
		}
	}

}
