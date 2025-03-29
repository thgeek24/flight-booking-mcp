package models

import "time"

type Flight struct {
	ID             string    `json:"id"`
	Departure      string    `json:"departure"`
	Destination    string    `json:"destination"`
	DepartureTime  time.Time `json:"departureTime"`
	ArrivalTime    time.Time `json:"arrivalTime"`
	Price          float64   `json:"price"`
	AvailableSeats int       `json:"availableSeats"`
}
