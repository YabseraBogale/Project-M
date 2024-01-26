package main

import (
	"io/ioutil"
	"path"
)

func main() {

	file, err := ioutil.ReadDir(".")
	if err != nil {

	}
	for _, i := range file {
		println(path.Dir(i.Name()))
	}

}
