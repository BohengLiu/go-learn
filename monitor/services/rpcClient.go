package services

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func InitClient() *ethclient.Client {
	client, err := ethclient.Dial("wss://eth-mainnet.g.alchemy.com/v2/kS8SZQU8zi48DdX5SIf8xMzAi6Imj7HZ")
	if err != nil {
		log.Fatal(err)
	}
	return client
}
