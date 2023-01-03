package main

import (
	"fmt"
	"helloWorld/hi"
)

func main() {
	fmt.Println("")

	s := hi.SayHi()
	fmt.Printf("s: %s\n", s)

	fmt.Println(s)

}
