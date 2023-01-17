package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var infuraURL = "https://mainnet.infura.io/v3/7fdd6b5a027641cf910c6c1cc6635610"

func main() {

	//Blockchain
	client, err := ethclient.DialContext(context.Background(), infuraURL)

	if err != nil {
		log.Fatal("Error to create clinet: ", err)
	}

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
