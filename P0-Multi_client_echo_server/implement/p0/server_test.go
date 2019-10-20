package p0

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	port := s.Count()
	fmt.Println(port)
}

func TestSlice(t *testing.T) {
	a := ";" + strconv.FormatInt(1000, 10)
	fmt.Println(a)
}
