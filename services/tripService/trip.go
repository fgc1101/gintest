package tripService

import (
	"gintest/models/trip"
)

// IsTransferByTripDetail 是否中转
func IsTransferByTripDetail(details []trip.Detail) int {

	if details == nil {
		return 0
	}

	var laiData []interface{}
	var fanData []interface{}

	for _, v := range details {
		if v.IsBack == 0 {
			laiData = append(laiData, v)
		} else {
			fanData = append(fanData, v)
		}
	}

	if len(laiData) > 1 {
		return 1

	}

	if len(fanData) > 1 {
		return 1
	}

	return 0
}

// 获取行程的住宿列表
