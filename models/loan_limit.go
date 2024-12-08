package models

import "gorm.io/gorm"

type LoanLimit struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"not null"`
	Tenor  int     `gorm:"not null"`
	Amount float64 `gorm:"not null"`
}

func MigrateLoanLimit(db *gorm.DB) {
	db.AutoMigrate(&LoanLimit{})
}
