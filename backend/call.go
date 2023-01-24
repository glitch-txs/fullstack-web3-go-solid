package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(a string, u string) (string, error) {

	infuraURL := "https://goerli.infura.io/v3/" + os.Getenv("INFURA_KEY")

	cl, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatal("Error to create clinet: ", err)
	}

	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		log.Fatal("Error to create clinet: ", err)
	}

	fmt.Println(chainid)

	defer cl.Close()

	ca := common.HexToAddress(a)

	cu := common.HexToAddress(u)

	ctr, err := NewERC20(ca, cl)

	o, err := ctr.BalanceOf(nil, cu)

	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
		return "", err
	}

	fbal := new(big.Float)
	fbal.SetString(o.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
	fmt.Println("Token name:", o)

	return value.String(), nil
}
