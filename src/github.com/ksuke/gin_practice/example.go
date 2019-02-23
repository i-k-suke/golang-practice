package main

import "github.com/gin-gonic/gin"

// terminal
// $ go run example.go

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
	//router.Run(":8080")	
}
