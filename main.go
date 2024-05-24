package main

import (
	"log"

	"dbo.api/controllers"
	"dbo.api/database"
	"dbo.api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create a new Gin router
	r := gin.Default()

	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		// Customer routes
		authorized.GET("/customers", controllers.GetCustomers)
		authorized.GET("/customers/:id", controllers.GetCustomer)
		authorized.POST("/customers", controllers.CreateCustomer)
		authorized.PUT("/customers/:id", controllers.UpdateCustomer)
		authorized.DELETE("/customers/:id", controllers.DeleteCustomer)
		authorized.GET("/customers/search", controllers.SearchCustomers)

		// Order routes
		authorized.GET("/orders", controllers.GetOrders)
		authorized.GET("/orders/:id", controllers.GetOrder)
		authorized.POST("/orders", controllers.CreateOrder)
		authorized.PUT("/orders/:id", controllers.UpdateOrder)
		authorized.DELETE("/orders/:id", controllers.DeleteOrder)
		authorized.GET("/orders/search", controllers.SearchOrders)
	}

	// Log that the server is starting
	log.Println("Starting server on port 8080")

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
