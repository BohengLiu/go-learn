package p0

import (
	"fmt"
	"net"
	"strconv"
)

type multiEchoServer struct {
	msgs []string
	conns []net.Conn
	msgChain chan string
	listener net.Listener
	count    int
	port     int
}

func New() MultiEchoServer {
	return &multiEchoServer{count: 0, port: 1}
}

func (s *multiEchoServer) Start(port int) error {
	// TODO: s
	s.conns = make([]net.Conn,100)
	ln, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(port),10))
	s.listener = ln
	fmt.Println("server start, is listening ", port)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err!= nil {
			continue
		}
		go s.handleRequest(conn)
	}
}

func (s *multiEchoServer) handleRequest(conn net.Conn) {
	for {
		buf := make([]byte,4096)
		n, err := conn.Read(buf)
		if err!= nil {
			fmt.Println(err)
			break
		}
		fmt.Println("msg length is",n)
		fmt.Println(string(buf[0:n]))
	}
	//buf.ReadString('\n')
}

func (s *multiEchoServer) addConnection(conn net.Conn) {
	s.conns = append(s.conns,conn)
}


func (s *multiEchoServer) Count() int {
	return s.count
}

func (s *multiEchoServer) Close() {
	return
}
