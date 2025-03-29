package handlers

import (
	"flight-booking-service/repository"
	"github.com/gin-gonic/gin"
)

func GetFlights(repo repository.FlightRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		flights, err := repo.GetAllFlights()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, flights)
	}
}
