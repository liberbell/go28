package main

import (
	"fmt"
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
		}
		io.WriteString(conn, "\nHello from server")
		fmt.Fprintln(conn, "\nHow is your day")
		fmt.Fprintf(conn, "")
	}
}
