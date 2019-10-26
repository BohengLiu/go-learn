package p0

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
)

type multiEchoServer struct {
	msgQ     *MsgQueue
	connMap  map[string]net.Conn
	msgChain chan int
	listener net.Listener
	//count    int
	//port int
}

func New() MultiEchoServer {
	return &multiEchoServer{}
}

func (s *multiEchoServer) Start(port int) error {
	// TODO:
	s.connMap = make(map[string]net.Conn)
	s.msgQ = &MsgQueue{}
	s.msgChain = make(chan int)
	ln, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(port), 10))
	s.listener = ln
	fmt.Println("server start, is listening ", port)
	if err != nil {
		return err
	}
	defer s.Close()
	go s.handleMsgQueueTask()
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go s.handleRequest(conn)
	}
}

func (s *multiEchoServer) handleRequest(conn net.Conn) {
	connId := conn.RemoteAddr().Network() + "/" + conn.RemoteAddr().String()
	fmt.Println(connId)
	s.addConnection(connId, conn)
	defer s.removeConnection(connId)
	for {
		rawMsg := make([]byte, 0)
		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println(connId + " disconnect!")
					return
				} else {
					fmt.Println(err)
					continue
				}

			}
			rawMsg = append(rawMsg, buf[0:n]...)
			if string(buf[n-1]) == "\n" {
				fmt.Println("content is:", string(rawMsg))
				go s.broadcastMsg(string(rawMsg))
				rawMsg = make([]byte, 0)
			}
		}
	}

}

func (s *multiEchoServer) addConnection(connId string, conn net.Conn) {
	s.connMap[connId] = conn
	//s.count = s.count + 1
}

func (s *multiEchoServer) removeConnection(connId string) {
	s.connMap[connId].Close()
	delete(s.connMap, connId)

}

func (s *multiEchoServer) Count() int {
	return len(s.connMap)
}

func (s *multiEchoServer) broadcastMsg(msg string) {
	for k, _ := range s.connMap {
		go s.sendMsg(k, msg)
		//byteMsg := []byte(msg)
		//pos := 0
		//for {
		//	n,err := v.Write(byteMsg[pos:])
		//	if err != nil {
		//		fmt.Println("broadcastMsg error:", err)
		//		go s.handleReSendMsg(k,msg)
		//	}
		//	if n == len([]byte(msg)) {
		//		break
		//	} else {
		//		pos = n
		//	}
		//}
	}
	//return
}

func (s *multiEchoServer) handleReSendMsg(connId, msg string) {
	s.msgQ.appendNode(connId, msg)
	s.msgChain <- '1'
}

func (s *multiEchoServer) handleMsgQueueTask() {
	for {
		for {
			if s.msgQ.Count > 0 {
				msgNode, err := s.msgQ.getNode()
				if err != nil {
					break
				}
				go s.sendMsg(msgNode.To, msgNode.Msg)
			}
		}

		select {
		case <-s.msgChain:

		}
	}
}

func (s *multiEchoServer) sendMsg(connId, msg string) {
	//defer s.removeConnection(connId)
	conn := s.connMap[connId]
	byteMsg := []byte(msg)
	pos := 0
	for {
		n, err := conn.Write(byteMsg[pos:])
		if err != nil {
			fmt.Println("broadcastMsg error:", err)
			go s.handleReSendMsg(connId, msg)
		}
		if n == len([]byte(msg)) {
			break
		} else {
			pos = n
		}
	}
}

type MsgNode struct {
	To   string
	Msg  string
	Next *MsgNode
}

type MsgQueue struct {
	First *MsgNode
	Count int
	Last  *MsgNode
}

func (q *MsgQueue) appendNode(to, msg string) {
	newNode := MsgNode{to, msg, nil}
	if q.Count > 0 {
		q.Last.Next = &newNode
		q.Count += 1
		q.Last = &newNode
	} else {
		q.First = &newNode
		q.Count = 1
		q.Last = &newNode
	}
}

func (q *MsgQueue) getNode() (*MsgNode, error) {
	if q.Count > 0 {
		result := q.First
		if q.First.Next != nil {
			q.First = q.First.Next
			q.Count -= 1
		} else {
			q.First = nil
			q.Last = nil
			q.Count = 0
		}
		return result, nil
	} else {
		return nil, errors.New("the queue is empty")
	}
}

func (s *multiEchoServer) Close() {
	for _, v := range s.connMap {
		v.Write([]byte("end"))
		v.Close()
	}
}
