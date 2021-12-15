package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	router := r.Group("/test")
	router2 := router.Group("/test")
	router2.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
