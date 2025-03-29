package handlers

import (
	"flight-booking-service/models"
	"flight-booking-service/repository"
	"github.com/gin-gonic/gin"
)

func BookFlight(flightRepo repository.FlightRepository, bookingRepo repository.BookingRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			FlightID string `json:"flightId"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		// Check if flight exists
		flight, err := flightRepo.GetFlightById(request.FlightID)
		if err != nil {
			c.JSON(404, gin.H{"error": "flight not found"})
			return
		}

		// Create booking
		booking := models.Booking{
			FlightID: flight.ID,
			Status:   "confirmed",
		}
		bookingID, err := bookingRepo.CreateBooking(booking)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to book flight"})
			return
		}

		c.JSON(201, gin.H{
			"bookingId": bookingID,
			"status":    "confirmed",
		})
	}
}
