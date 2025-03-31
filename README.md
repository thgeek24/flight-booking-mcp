# Flight Booking Service

A RESTful API service for flight booking management built with Go and Gin framework.

## Project Structure

```
flight-booking-service/
├── handlers/
│   ├── flight_handler.go
│   └── booking_handler.go
├── models/
│   ├── flight.go
│   └── booking.go
├── repository/
│   ├── flight_repo.go
│   └── booking_repo.go
├── main.go
└── README.md
```

## API Endpoints

- `GET /flights` - List all available flights
- `POST /bookings` - Create a new booking
- `GET /bookings` - List all bookings

## Setup

1. Clone the repository
2. Install dependencies:
```bash
go mod download
```
3. Run the server:
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## Development

- Frontend development server runs at `http://localhost:5173`
- CORS is configured to allow requests from the frontend
- Using in-memory repositories for data storage

## Technologies

- Go 1.x
- Gin Web Framework
- CORS middleware