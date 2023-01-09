package tasks

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeBlock(client *ethclient.Client) {

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())      // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64()) // 3477413
			fmt.Println(block.Time())            // 1529525947
			for _, tx := range block.Transactions() {

				v := tx.Value()
				fbalance := new(big.Float)
				fbalance.SetString(v.String())
				ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
				if ethValue.Cmp(big.NewFloat(1)) > 0 {
					fmt.Println("is bigger than 1 ether")
				}

				fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
				fmt.Println(ethValue)        // 10000000000000000
				fmt.Println(tx.To())
			}
		}
	}
}
