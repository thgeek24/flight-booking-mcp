package main

import (
	"flight-booking-service/handlers"
	"flight-booking-service/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	flightRepo := repository.NewInMemoryFlightRepo()
	bookingRepo := repository.NewInMemoryBookingRepo()

	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vite's default port
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/flights", handlers.GetFlights(flightRepo))
	r.POST("/bookings", handlers.BookFlight(flightRepo, bookingRepo))
	r.GET("/bookings", handlers.GetBookings(bookingRepo))

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
