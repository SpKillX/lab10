package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})
	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
