package repository

import (
	"errors"
	"flight-booking-service/models"
	"time"
)

type FlightRepository interface {
	GetAllFlights() ([]models.Flight, error)
	GetFlightById(id string) (models.Flight, error)
	BookSeat(flightId string) error
}

type InMemoryFlightRepo struct {
	flights []models.Flight
}

func NewInMemoryFlightRepo() *InMemoryFlightRepo {
	return &InMemoryFlightRepo{
		flights: []models.Flight{
			{
				ID:             "1",
				Departure:      "New York",
				Destination:    "San Francisco",
				DepartureTime:  must(time.Parse(time.RFC3339, "2025-03-29T10:10:00Z")),
				ArrivalTime:    must(time.Parse(time.RFC3339, "2025-03-29T13:15:00Z")),
				Price:          299.99,
				AvailableSeats: 5,
			},
			{
				ID:             "2",
				Departure:      "San Francisco",
				Destination:    "New York",
				DepartureTime:  must(time.Parse(time.RFC3339, "2025-03-30T09:15:00Z")),
				ArrivalTime:    must(time.Parse(time.RFC3339, "2025-03-30T17:00:00Z")),
				Price:          399.99,
				AvailableSeats: 10,
			},
			{
				ID:             "3",
				Departure:      "Los Angeles",
				Destination:    "Chicago",
				DepartureTime:  must(time.Parse(time.RFC3339, "2025-03-31T08:00:00Z")),
				ArrivalTime:    must(time.Parse(time.RFC3339, "2025-03-31T14:10:00Z")),
				Price:          349.99,
				AvailableSeats: 8},
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

func (r *InMemoryFlightRepo) BookSeat(flightId string) error {
	for i, flight := range r.flights {
		if flightId == flight.ID {
			if flight.AvailableSeats > 0 {
				r.flights[i].AvailableSeats--
				return nil
			}
			return errors.New("no available seats")
		}
	}
	return errors.New("flight not found")
}

func must(t time.Time, err error) time.Time {
	if err != nil {
		panic(err)
	}
	return t
}
