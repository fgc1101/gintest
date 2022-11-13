package trip

import (
	"gintest/models/common"
	"log"
)

type Dorm struct {
	DormId        int         `json:"dorm_id" gorm:"type:int(11)"`
	TripId        int         `json:"trip_id" gorm:"type:int(11)"`
	StartTime     string      `json:"start_time" gorm:"type:text"`
	EndTime       string      `json:"end_time" gorm:"type:text"`
	DormDays      int         `json:"dorm_days" gorm:"type:int(11)"`
	Address       string      `json:"address" gorm:"type:text"`
	ReceiptGiver  string      `json:"receipt_giver" gorm:"type:text"`
	ReceiptImages interface{} `json:"receipt_images" gorm:"type:text"`
	ReceiptMoney  float64     `json:"receipt_money" gorm:"type:float(10,2)"`
	CreatedAt     string      `json:"created_at" gorm:"type:text"`
	UpdatedAt     string      `json:"updated_at" gorm:"type:text"`
}

func (t *Dorm) TableName() string {
	return "as_market_trip_dorm"
}

func (t *Dorm) GetDormRow(tripId uint) *Dorm {
	var DormRow *Dorm
	err := common.DB.Model(&Dorm{}).
		Debug().
		Where("trip_id = ?", tripId).
		Order("CreatedAt desc").
		Limit(1).
		First(&DormRow).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return DormRow
}
