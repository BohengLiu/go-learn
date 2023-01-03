package main

import (
	"fmt"
)

func main() {
	fmt.Println("hhh")
	var a string = "aabbcc"
	var b *string = &a

	fmt.Println(*b)

}
