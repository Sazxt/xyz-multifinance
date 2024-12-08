package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if dbHost == "" {
		dbHost = "127.0.0.1"
		dbPort = "3306"
	}
	dsn := "root:123@tcp(" + dbHost + ":" + dbPort + ")/xyz_multifinance?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("Failed to set connection pool:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	DB = database
}
