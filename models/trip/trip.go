package trip

import (
	"gintest/models/common"
	"gorm.io/gorm"
	"time"
)

type Trip struct {
	Id         uint        `json:"id" gorm:"type:int(11)"`
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
	TripStatus int         `json:"trip_status" gorm:""`
	CreatedAt  time.Time   `json:"created_at" gorm:""`
	UpdatedAt  time.Time   `json:"updated_at" gorm:""`
	DeletedAt  time.Time   `json:"deleted_at"" gorm:"index;default:null"`
	Details    []Detail    `json:"details" gorm:"foreignkey:TripId"`
	Dorm       Dorm        `json:"dorm"`
}

func (table *Trip) TableName() string {
	return "as_market_trip"
}

func GetTripList(where map[string]interface{}) *gorm.DB {
	cond, vals, _ := common.WhereBuild(where)
	return common.DB.Model(&Trip{}).Debug().Order("created_at desc").Where(cond, vals...)

}
