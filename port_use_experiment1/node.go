package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("ok")
	ln, err := net.Listen("tcp",":9090")
	if err != nil {
		return
	}
	go callserver()
	//select {}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	for {
		rawMsg := make([]byte, 0)
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println(  " disconnect!")
					return
				} else {
					fmt.Println(err)
					continue
				}

			}
			rawMsg = append(rawMsg, buf[0:n]...)
			if n < 4096 {
				fmt.Println("content is:", string(rawMsg))
				rawMsg = make([]byte, 0)
			}
		}
	}
}

func callserver() {
	for {
		time.Sleep(5*time.Second)
		lnetAddr := &net.TCPAddr{Port:9090}
		rnetAddr := &net.TCPAddr{Port:10001}
		//d := net.Dialer{Timeout: time.Second*5,LocalAddr: netAddr}
		socket,err := net.DialTCP("tcp",lnetAddr,rnetAddr)
		fmt.Println(socket)
		//defer socket.Close()
		if err!=nil {
			fmt.Println(err)
			continue
		}
		str := []byte("hello")
		_, err = socket.Write(str)
		if err != nil {
			fmt.Println(err)
		}
		socket.Close()

	}

}

