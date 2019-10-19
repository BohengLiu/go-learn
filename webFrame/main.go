package main

import (
	"fmt"
	"net"
)

func recvFile(conn net.Conn) {
	fmt.Println("new connection: ", conn.RemoteAddr())
	b := make([]byte)
	conn.Read(b)
	fmt.Println(string(b))
	defer conn.Close()
}

func main() {
	fmt.Println("start listening :8080")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go recvFile(conn)
	}
}
