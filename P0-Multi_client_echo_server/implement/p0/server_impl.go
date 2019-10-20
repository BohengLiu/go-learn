package p0

import (
	"fmt"
	"net"
)

type multiEchoServer struct {
	listener net.Listener
	count    int
	port     int
}

func New() MultiEchoServer {
	return &multiEchoServer{count: 0, port: 1}
}

func (s *multiEchoServer) Start(port int) error {
	// TODO: s
	var err error
	s.listener, err = net.Listen("tcp", ":"+string(port))
	fmt.Println("server start, is listening ", port)
	if err != nil {
		return err
	}
	return nil
}

func (s *multiEchoServer) Count() int {
	return s.count
}

func (s *multiEchoServer) Close() {
	return
}
