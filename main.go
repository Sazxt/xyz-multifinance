package main

import (
	"xyz-multifinance/config"
	"xyz-multifinance/controllers"
	"xyz-multifinance/middleware"
	"xyz-multifinance/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	// Migrate models
	models.MigrateUser(config.DB)
	models.MigrateLoanLimit(config.DB)
	models.MigrateTransaction(config.DB)

	r := gin.Default()
	r.Use(middleware.SecureHeaders())

	// Routes
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.POST("/transactions", controllers.CreateTransaction)

	// Run server
	r.Run(":8080")
}
