// File: internal/api/router.go

package api

import (
	"glofox-backend/internal/api/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter configures all the routes and handlers
func SetupRouter(classHandler *handlers.ClassHandler, bookingHandler *handlers.BookingHandler) *gin.Engine {
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Home endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":       "Welcome to Glofox Studio API",
			"documentation": "/swagger/index.html",
		})
	})

	// Class routes
	classRoutes := router.Group("/classes")
	{
		classRoutes.POST("/", classHandler.CreateClass)
		classRoutes.GET("/", classHandler.GetAllClasses)
		classRoutes.GET("/:id", classHandler.GetClassByID)
	}

	// Booking routes
	bookingRoutes := router.Group("/bookings")
	{
		bookingRoutes.POST("/", bookingHandler.CreateBooking)
		bookingRoutes.GET("/", bookingHandler.GetAllBookings)
		bookingRoutes.GET("/:id", bookingHandler.GetBookingByID)
	}

	return router
}
