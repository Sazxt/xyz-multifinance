package controllers

import (
	"context"
	"fmt"
	"net/http"
	// "sync"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
	"time"

	"gorm.io/gorm"
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


	var user models.User
	if err := config.DB.First(&user, input.UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var loanLimit models.LoanLimit
	if err := config.DB.Where("user_id = ? AND tenor = ?", input.UserID, input.Installment).First(&loanLimit).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Loan limit not available for this tenor"})
		return
	}

	if input.OTR > loanLimit.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient loan limit"})
		return
	}

	loanLimit.Amount -= input.OTR
	config.DB.Save(&loanLimit)

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	auditLog := models.AuditLog{
		UserID:    user.ID,
		Action:    "CREATE_TRANSACTION",
		Details:   "Transaction created with OTR: " + fmt.Sprintf("%.2f", input.OTR),
		Timestamp: time.Now(),
	}
	config.DB.Create(&auditLog)

	c.JSON(http.StatusOK, gin.H{
		"transaction": input,
		"message":     "Transaction successfully created, remaining limit updated",
	})
}
