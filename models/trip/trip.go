package trip

import "gorm.io/gorm"

type Trip struct {
	gorm.Model
	TripNumber string      `json:"trip_number"`
	SchoolID   int         `json:"school_id"`
	TripTypes  int         `json:"trip_types"`
	TripReason string      `json:"trip_reason"`
	TripMode   string      `json:"trip_mode"`
	IsDorm     string      `json:"is_dorm"`
	IsStay     string      `json:"is_stay"`
	IsCounting string      `json:"is_counting"`
	Attachment interface{} `json:"attachment"`
	AdminID    int         `json:"admin_id"`
	TripStatus int         `json:"trip_status"`
}

func (table *Trip) TableName() string {
	return "as_market_trip"
}
