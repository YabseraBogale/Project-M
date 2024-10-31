package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		log.Println(err)
	}
	var input string
	for _, j := range file {

		fmt.Println(file[j])
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
	}
	// scanner := bufio.NewScanner(bytes.NewBuffer(file))
	// count := 0
	// slower := 1
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	slower += 1

	// 	if strings.Contains(line, "EmailAddress") {
	// 		count += 1
	// 		fmt.Println(line)
	// 	}
	// 	if slower%100 == 0 {
	// 		println("At line", slower)
	// 		time.Sleep(30 * time.Second)
	// 	}
	// 	if slower == 10000 {
	// 		break
	// 	}

	// }
	// println("found", count, "email")
}
