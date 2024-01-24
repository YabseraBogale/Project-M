package main

import "io/ioutil"

func main() {

	file, err := ioutil.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		println(i.Mode())
	}

}
