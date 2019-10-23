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
	defer s.Close()
	for {
		conn, err := ln.Accept()
		if err!= nil {
			continue
		}
		go s.handleRequest(conn)
	}
}

func (s *multiEchoServer) handleRequest(conn net.Conn) {
	// TODO: go 的队列维护
	s.addConnection(conn)
	for {
		buf := make([]byte,4096)
		n, err := conn.Read(buf)
		if err!= nil {
			fmt.Println(err)
			break
		}
		fmt.Println("msg length is",n)
		fmt.Println("content is:", string(buf[0:n]))
		conn.Write([]byte("receive"))
	}
	//buf.ReadString('\n')
	defer conn.Close()
}

func (s *multiEchoServer) addConnection(conn net.Conn) {
	s.conns = append(s.conns,conn)
	s.count = s.count + 1
}


func (s *multiEchoServer) Count() int {
	return s.count
}

func (s *multiEchoServer) Close() {
	for i:=0;i<len(s.conns);i++ {
		s.conns[i].Write([]byte("end"))
		s.conns[i].Close()
	}
	return
}

