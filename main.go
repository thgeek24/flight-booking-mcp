package main

import (
	"flight-booking-service/handlers"
	"flight-booking-service/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	flightRepo := repository.NewInMemoryFlightRepo()
	bookingRepo := repository.NewInMemoryBookingRepo()

	r := gin.Default()

	r.GET("/flights", handlers.GetFlights(flightRepo))
	r.POST("/bookings", handlers.BookFlight(flightRepo, bookingRepo))
	r.GET("/bookings", handlers.GetBookings(bookingRepo))

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
