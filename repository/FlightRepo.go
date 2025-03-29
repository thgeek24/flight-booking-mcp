package repository

import (
	"errors"
	"flight-booking-service/models"
	"time"
)

type FlightRepository interface {
	GetAllFlights() ([]models.Flight, error)
	GetFlightById(id string) (models.Flight, error)
}

type InMemoryFlightRepo struct {
	flights []models.Flight
}

func NewInMemoryFlightRepo() *InMemoryFlightRepo {
	return &InMemoryFlightRepo{
		flights: []models.Flight{
			{
				ID:            "1",
				Departure:     "NYC",
				Destination:   "LAX",
				DepartureTime: must(time.Parse(time.RFC3339, "2025-03-29T10:00:00Z")),
				ArrivalTime:   must(time.Parse(time.RFC3339, "2023-03-29T13:00:00Z")),
				Price:         299.99,
			},
			{
				ID:            "2",
				Departure:     "SFO",
				Destination:   "JFK",
				DepartureTime: must(time.Parse(time.RFC3339, "2023-03-30T09:00:00Z")),
				ArrivalTime:   must(time.Parse(time.RFC3339, "2023-03-30T17:00:00Z")),
				Price:         399.99,
			},
		},
	}
}

func (r *InMemoryFlightRepo) GetAllFlights() ([]models.Flight, error) {
	return r.flights, nil
}

func (r *InMemoryFlightRepo) GetFlightById(id string) (models.Flight, error) {
	for _, f := range r.flights {
		if f.ID == id {
			return f, nil
		}
	}
	return models.Flight{}, errors.New("flight not found")
}

func must(t time.Time, err error) time.Time {
	if err != nil {
		panic(err)
	}
	return t
}
