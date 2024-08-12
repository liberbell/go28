package main

import (
	"io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "Hello, this is a server")
		conn.Close()
	}
}
