package models

import "gorm.io/gorm"

type Transaction struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	ContractNo  string  `gorm:"unique;not null"`
	OTR         float64 `gorm:"not null"`
	AdminFee    float64
	Installment float64
	Interest    float64
	AssetName   string
}

func MigrateTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
