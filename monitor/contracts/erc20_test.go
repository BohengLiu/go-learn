package contracts

import (
	"fmt"
	"testing"
)

func TestAbi(t *testing.T) {
	contractAbi := GetErc20Abi()
	// events := make(map[string]*abi.Event)
	for idx, val := range contractAbi.Events {
		// events[val.ID.String()]
		fmt.Println(idx)
		fmt.Println(val.ID.String())
	}
}
