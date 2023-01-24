package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//server
	r := gin.Default()

	r.GET("/balance", func(c *gin.Context) {

		ca := c.Query("contract")

		ua := c.Query("user")

		fmt.Printf("%v and %v /n", ca, ua)

		b, err := GetBalance(ca, ua)

		if err != nil {
			c.JSON(400, gin.H{
				"message": "something went wrong >.<'",
			})
		}
		c.JSON(200, gin.H{
			"message": b,
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
