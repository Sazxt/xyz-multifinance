package routes

import (
	"xyz-multifinance/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)
	r.POST("/transactions", controllers.CreateTransaction)
}
