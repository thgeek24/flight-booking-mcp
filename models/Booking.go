package models

type Booking struct {
	ID       string `json:"id"`
	FlightID string `json:"flightID"`
	Status   string `json:"status"`
}
