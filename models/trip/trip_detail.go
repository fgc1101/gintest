package trip

type Detail struct {
	DetailId              int     `json:"detail_id" gorm:"type:int(11)"`
	TripId                int     `json:"trip_id" gorm:"type:int(11)"`
	DetailType            int     `json:"detail_type" gorm:"type:int(11)"`
	IsBack                int     `json:"is_back" gorm:"type:int(11)"`
	LocationStartPlace    string  `json:"location_start_place" gorm:"type:text"`
	LocationStartPosition string  `json:"location_start_position" gorm:"type:text"`
	LocationEndPlace      string  `json:"location_end_place" gorm:"type:text"`
	LocationEndPosition   string  `json:"location_end_position" gorm:"type:text"`
	StartPlace            string  `json:"start_place" gorm:"type:text"`
	EndPlace              string  `json:"end_place" gorm:"type:text"`
	StartPosition         string  `json:"start_position" gorm:"type:text"`
	EndPosition           string  `json:"end_position" gorm:"type:text"`
	StartTime             string  `json:"start_time" gorm:"type:text"`
	EndTime               string  `json:"end_time" gorm:"type:text"`
	Distance              float64 `json:"distance" gorm:"type:float(10,2)"`
	LocationDistance      float64 `json:"location_distance" gorm:"type:float(10,2)"`
	DetailMoney           float64 `json:"detail_money" gorm:"type:float(10,2)"`
	DetailStatus          int     `json:"detail_status" gorm:"type:int(11)"`
	ReimbursementInfoId   int     `json:"reimbursement_info_id" gorm:"type:int(11)"`
	Prove                 string  `json:"prove" gorm:"type:text"`
	ElcSheetUrl           string  `json:"elc_sheet_url" gorm:"type:text"`
	CreatedAt             string  `json:"created_at" gorm:"type:text"`
	UpdatedAt             string  `json:"updated_at" gorm:"type:text"`
}

func (t *Detail) TableName() string {
	return "as_market_trip_detail"
}
