package main

import (
	"fmt"
	"github.com/BohengLiu/go-learn/P0-Multi_client_echo_server/implement/p0"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("start p0")
	server := p0.New()
	go monitor(server)
	err := server.Start(8090)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func monitor(server p0.MultiEchoServer) {
	for {
		fmt.Println("Current number of connections is: " + strconv.FormatInt(int64(server.Count()), 10))
		time.Sleep(time.Second * 20)
	}
}
