package main

import (
	"fmt"
	"os"
)

func main() {
	print("Enter File To Search: ")
	var input string
	fmt.Scanln(&input)
	fileNAME := "./" + input
	fileinfo, err := os.Stat(fileNAME)
	if err != nil {

	}
	println(fileinfo)

}
