package bitcoin

import "testing"
import "fmt"

func Test_hash(t *testing.T) {
	result := Hash("hello", 12345)
	fmt.Println(result)
}