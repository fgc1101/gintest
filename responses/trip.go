package responses

import "gintest/models/trip"

type TripListResponse struct {
	*trip.Trip
	IsTransfer int `json:"is_transfer"`
	Dorm       *trip.Dorm
}
