package repository

import (
	"flight-booking-service/models"
	"github.com/google/uuid"
)

type BookingRepository interface {
	CreateBooking(booking models.Booking) (string, error)
}

type InMemoryBookingRepo struct {
	bookings []models.Booking
}

func NewInMemoryBookingRepo() *InMemoryBookingRepo {
	return &InMemoryBookingRepo{
		bookings: []models.Booking{},
	}
}

func (r *InMemoryBookingRepo) CreateBooking(booking models.Booking) (string, error) {
	booking.ID = uuid.New().String()
	r.bookings = append(r.bookings, booking)
	return booking.ID, nil
}
