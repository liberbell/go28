package main

import (
	"bufio"
	"fmt"
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
		scanner := bufio.NewScanner()
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
	}

}
