package main

import "fmt"

func printStr(str *string) {
	*str = `world`
}

func main() {
	// s := `hello`
	// s2 := string([]byte{s[1], s[3]})
	// // printStr(&s)

	// fmt.Println(s2)
	s3 := []byte{}
	s3 = append(s3, 0x44, 0x66)
	fmt.Println(string(s3))
}
