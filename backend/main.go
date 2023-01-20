package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	infuraURL := "https://mainnet.infura.io/v3/" + os.Getenv("INFURA_KEY")

	fmt.Println(infuraURL)

	//Blockchain
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

	//server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
