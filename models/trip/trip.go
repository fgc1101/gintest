package trip

import (
	"gintest/models/common"
	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	TripNumber string      `json:"trip_number" gorm:"type:text"`
	SchoolID   int         `json:"school_id" gorm:"type:int(11)"`
	TripTypes  int         `json:"trip_types" gorm:"type:int(11)"`
	TripReason string      `json:"trip_reason" gorm:"type:text"`
	TripMode   string      `json:"trip_mode" gorm:"type:text"`
	IsDorm     string      `json:"is_dorm" gorm:"type:text"`
	IsStay     string      `json:"is_stay" gorm:"type:text"`
	IsCounting string      `json:"is_counting" gorm:"type:text"`
	Attachment interface{} `json:"attachment" gorm:"type:text"`
	AdminID    int         `json:"admin_id" gorm:"type:int(11)"`
	TripStatus int         `json:"trip_status" gorm:"type:int(11)"`
}

func (table *Trip) TableName() string {
	return "as_market_trip"
}

func GetTripList(where map[string]interface{}) *gorm.DB {
	cond, vals, _ := common.WhereBuild(where)
	return common.DB.Model(new(Trip)).Debug().Where(cond, vals...)

}
