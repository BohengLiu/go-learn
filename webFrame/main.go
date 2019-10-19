package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn) {
	fmt.Println("new connection: ", conn.RemoteAddr())
	result := bytes.NewBuffer(nil)
	for {
		var buf [10]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			panic(err)
		}
		result.Write(buf[0:n])
		fmt.Println(buf[0:n], n)
	}

	fmt.Println(string(result.Bytes()))
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
		go handler(conn)
	}
}
