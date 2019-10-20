package main

import (
	"fmt"
	"github.com/BohengLiu/go-learn/P0-Multi_client_echo_server/implement/p0"
)

func main() {
	fmt.Println("start p0")
	server := p0.New()
	count := server.Count()
	fmt.Println(count)
}
