package main

import (
	"bufio"
	"net"
)

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
}
