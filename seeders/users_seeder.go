package seeders

import (
	"log"

	"xyz-multifinance/models"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{ID: 1, NIK: "119283900018391", FullName: "Aldi", LegalName: "Aldi", BirthPlace: "Tanggerang", BirthDate: "16-03-2000", Salary: 5000000, PhotoKTP: "none", PhotoSelfie: "none", LoanLimits: []models.LoanLimit{
			{UserID: 1, Tenor: 1, Amount: 100000},
			{UserID: 1, Tenor: 2, Amount: 200000},
			{UserID: 1, Tenor: 3, Amount: 500000},
			{UserID: 1, Tenor: 6, Amount: 700000},
		}},
	}
	for _, user := range users {
		var existingUser models.User
		result := db.Where("nik = ? AND id = ?", user.NIK, user.ID).First(&existingUser)
		if result.Error != nil {
			if err := db.Create(&user).Error; err != nil {
				log.Printf("Failed to seed user : %v", err)
			} else {
				log.Printf("Created user name %s, NIK %s", user.FullName, user.NIK)
			}
		} else {
			log.Printf("User name %s, NIK %s already exists, Skipping.", user.FullName, user.NIK)
		}
	}
}
