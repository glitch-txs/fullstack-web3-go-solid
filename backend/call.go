package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CallContract(addr string) {

	infuraURL := "https://goerli.infura.io/v3/" + os.Getenv("INFURA_KEY")

	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatal("Error to create clinet: ", err)
	}

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Error to create clinet: ", err)
	}

	fmt.Println(chainid)

	defer client.Close()

	newAddr := common.HexToAddress(addr)

	ctr, err := NewERC20(newAddr, client)

	output, err := ctr.Name(nil)

	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", output)
}
