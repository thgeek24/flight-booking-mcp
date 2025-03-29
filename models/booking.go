package models

type Booking struct {
	ID       string `json:"id"`
	FlightID string `json:"flightId"`
	Status   string `json:"status"`
}
