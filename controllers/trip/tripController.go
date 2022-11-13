package trip

import (
	"gintest/define"
	"gintest/models/trip"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetTripList
// @Summary 获取行程列表
// @Schemes
// @Description 获取行程列表
// @Tags 行程
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "trip_number"
// @Success 200 {string} json "{"code": 200, "msg":"", "data":""}"
// @Router /trip/list [get]
func GetTripList(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultLimit))
	keyword := c.Query("keyword")
	offset := (page - 1) * size
	var totalNum int64

	var tripList []*trip.Trip

	where := map[string]interface{}{
		"trip_number like": "%" + keyword + "%",
		"admin_id in":      []int{27, 46},
	}

	err := trip.GetTripList(where).
		Count(&totalNum).
		Offset(offset).
		Limit(size).
		Preload("Details").
		Find(&tripList).Error
	if err != nil {
		log.Println(err)
		return
	}

	data := map[string]interface{}{
		"total": totalNum,
		"list":  tripList,
	}
	c.JSON(200, gin.H{"code": 200, "data": data, "msg": "成功"})

}
