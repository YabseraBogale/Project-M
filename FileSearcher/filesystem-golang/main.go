package main

import (
	"io/ioutil"
	"path"
)

func main() {

	file, err := ioutil.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {

		ff, err := ioutil.ReadDir(path.Dir(i.Name()))
		if err != nil {

		}
		for _, j := range ff {
			println(j.Name())
		}
	}

}
