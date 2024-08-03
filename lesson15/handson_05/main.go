package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

	}
}

func handle(conn net.Conn) {
	scnner := bufio.NewScanner(conn)
}
