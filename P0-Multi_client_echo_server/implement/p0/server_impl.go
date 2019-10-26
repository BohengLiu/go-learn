package p0

import (
	"fmt"
	"net"
	"strconv"
)

type multiEchoServer struct {
	msgs []string
	conns map[string]net.Conn
	msgChain chan string
	listener net.Listener
	//count    int
	port     int
}

func New() MultiEchoServer {
	return &multiEchoServer{port: 1}
}

func (s *multiEchoServer) Start(port int) error {
	// TODO: s
	s.conns = make(map[string]net.Conn)
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

	connId := conn.RemoteAddr().Network() + "/" + conn.RemoteAddr().String()
	fmt.Println(connId)
	s.addConnection(connId, conn)

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
	defer s.removeConnection(connId, conn)
}


func (s *multiEchoServer) addConnection(connId string, conn net.Conn) {
	s.conns[connId] = conn
	//s.count = s.count + 1
}

func (s *multiEchoServer) removeConnection(connId string, conn net.Conn) {
	delete(s.conns,connId)
	conn.Close()
}

func (s *multiEchoServer) Count() int {
	return len(s.conns)
}

func (s *multiEchoServer) broadcastMsg() {
	return
}

func (s *multiEchoServer) handleReSendMsg() {}


func (s *multiEchoServer) Close() {
	for _,v:= range s.conns {
		v.Write([]byte("end"))
		v.Close()
	}
}

