package seeders

import (
	"log"

	_ "fmt"
	"gorm.io/gorm"
	"xyz-multifinance/models"
)

func SeedLoanLimits(db *gorm.DB) {
	loanLimits := []models.LoanLimit{
		{UserID: 9, Tenor: 1, Amount: 100000},
		{UserID: 9, Tenor: 2, Amount: 200000},
		{UserID: 9, Tenor: 3, Amount: 500000},
		{UserID: 9, Tenor: 6, Amount: 700000},
	}

	for _, limit := range loanLimits {
		var existingLoanLimit models.LoanLimit
		result  := db.Where("user_id = ? AND tenor = ?", limit.UserID, limit.Tenor).First(&existingLoanLimit);
		if result.Error != nil {
			if err := db.Create(&limit).Error; err != nil {
				log.Printf("Failed to seed loan limit: %v", err)
			} else {
				log.Printf("Created loan limit for User %d, Tenor %d", limit.UserID, limit.Tenor)
			}
		} else {
			log.Printf("Loan limit for User %d, Tenor %d already exists. Skipping.", limit.UserID, limit.Tenor)
		}
	}
}
