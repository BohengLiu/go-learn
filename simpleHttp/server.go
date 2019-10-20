package main

import (
	"fmt"
	"io"
	"net"
)

func CreateServer() {
	fmt.Println("server start")
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("tcp server err:", err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			// 处理错误
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println(`get a conn`)
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			conn.Close()
		}
	}
	fmt.Println(string(buf[0:n]))
}
