package models

import (
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Action    string    `gorm:"not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`
	Details   string
}

func MigrateAuditLog(db *gorm.DB) {
	db.AutoMigrate(&AuditLog{})
}
