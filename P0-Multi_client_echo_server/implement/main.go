package main

import (
	"fmt"
	"github.com/BohengLiu/go-learn/P0-Multi_client_echo_server/implement/p0"
	"os"
)

func main() {
	fmt.Println("start p0")
	server := p0.New()
	err := server.Start(8090)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
