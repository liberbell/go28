package main

import (
	"fmt"
	"net"
)

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><title></title></head><body><strong>Hello world</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "content-length: %d\r\n", len(body))
	fmt.Fprint(conn, "content-type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
