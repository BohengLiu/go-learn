package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start the http")
	task := "server"
	if len(os.Args) > 1 {
		//for i:=0;i<len(os.Args);i++ {
		//	fmt.Println(os.Args[i])
		//}
		task = os.Args[1]
	}
	switch task {
	case "client":
		CreateClient()
	case "server":
		CreateServer()
	default:
		CreateServer()
	}
}
