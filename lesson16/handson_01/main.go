package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			fmt.Println("***Method", m)
		}
		if ln == "" {
			break
		}
		i++
	}
}
