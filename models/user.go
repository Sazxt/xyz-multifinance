package models

import "gorm.io/gorm"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	NIK          string `gorm:"unique;not null"`
	FullName     string
	LegalName    string
	BirthPlace   string
	BirthDate    string
	Salary       float64
	PhotoKTP     string
	PhotoSelfie  string
	LoanLimits   []LoanLimit   `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
