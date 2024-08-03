package main

import (
	"bufio"
	"log"
	"net"
	"strings"
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
	for scnner.Scan() {
		ln := strings.ToLower(scnner.Text())
		bs := []byte(ln)
	}
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
}
