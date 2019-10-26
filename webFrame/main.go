package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BohengLiu/go-learn/P0-Multi_client_echo_server/implement/p0"
)

//func handler(conn net.Conn) {
//	fmt.Println("new connection: ", conn.RemoteAddr())
//	result := bytes.NewBuffer(nil)
//	for {
//		var buf [10]byte
//		n, err := conn.Read(buf[0:])
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			fmt.Println(err)
//			panic(err)
//		}
//		result.Write(buf[0:n])
//		fmt.Println(buf[0:n], n)
//	}
//
//	fmt.Println(string(result.Bytes()))
//	defer conn.Close()
//}

func main() {
	fmt.Println("start listening :8080")
	server1 := p0.New()
	go monitor(server1)

	server1.Start(8080)
}

func monitor(s p0.MultiEchoServer) {
	for {
		time.Sleep(time.Second *10)
		fmt.Println("current connect", strconv.FormatInt(int64(s.Count()),10))
	}
}