package controllers

import (
	"context"
	// "fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
	"xyz-multifinance/config"
	"xyz-multifinance/models"
)

var ctx = context.Background()

func CreateTransaction(c *gin.Context) {
	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	var user models.User
	config.DB.First(&user, input.UserID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Proses transaksi
	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"transaction": input})
}
