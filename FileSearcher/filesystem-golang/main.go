package main

import "io/ioutil"

func main() {
	file, err := ioutil.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		List((i.Name()))
	}

}

func List(name string) string {
	if name == "done" {
		return "done"
	}
	file, err := ioutil.ReadDir(name)
	if err != nil {
		println("Error")
	}

	for _, i := range file {
		if i.IsDir() == true {
			println(i.Name())
			return List(i.Name())
		}
	}
	return "done"
}
