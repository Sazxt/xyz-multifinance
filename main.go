package main

import (
	"fmt"
	"xyz-multifinance/config"
	"xyz-multifinance/middleware"
	"xyz-multifinance/models"
	"xyz-multifinance/routes"
	"xyz-multifinance/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Migrate models
	models.MigrateUser(config.DB)
	models.MigrateLoanLimit(config.DB)
	models.MigrateTransaction(config.DB)
	models.MigrateAuditLog(config.DB)

	// Seed data
	// seeders.SeedLoanLimits(config.DB)
	seeders.SeedUsers(config.DB)

	r := gin.Default()
	r.Use(middleware.SecureHeaders())

	// Routes
	routes.RegisterRoutes(r)

	// Run server
	fmt.Println("Server running on port 8080")
	r.Run(":8080")
}
