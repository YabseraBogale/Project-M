package main

import "io/ioutil"

func main() {
	file, err := ioutil.ReadDir("./Folder")
	if err != nil {

	}
	for _, i := range file {
		println(i.Name())
		println(List((i.Name())))
	}

}

func List(name string) string {
	if name == "done" {
		println(name)
		return "done"
	}
	file, err := ioutil.ReadDir(name)
	if err != nil {
		println("Error")
	}

	for _, i := range file {
		if i.IsDir() == true {
			println(i.Name())
			return List("./" + i.Name())
		}
	}
	println(name)
	return "done"
}
