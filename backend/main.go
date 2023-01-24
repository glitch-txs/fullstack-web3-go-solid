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

	r.Use(CORSMiddleware())

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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
