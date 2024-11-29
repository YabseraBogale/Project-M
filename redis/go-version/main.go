package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Println(err)
		return
	}
	connection, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}
	defer connection.Close()
	for {
		buf := make([]byte, 1024)
		_, err := connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			os.Exit(1)

		}
	}
}
