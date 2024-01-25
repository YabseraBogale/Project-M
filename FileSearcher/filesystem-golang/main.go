package main

import "io/ioutil"

func main() {
	
}

func List(name string) bool{
	file,err:=ioutil.ReadDir(name)
	if err!=nil{
		println("Error")
	}
	
	for _,i:=range file{
		if(){

		}
	}
}