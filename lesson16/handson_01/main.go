package main

import "net"

func handle(conn net.Conn) {
	defer conn.Close()
}
