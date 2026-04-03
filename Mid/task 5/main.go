package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingRequest struct {
	UserID  int    `json:"user_id" binding:"required"`
	TableID int    `json:"table_id" binding:"required,gt=0"`
	Email   string `json:"email" binding:"required,email"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/book", func(c *gin.Context) {
		var json BookingRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Booking confirmed", "data": json})
	})

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}
