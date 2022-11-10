package trip

import (
	"fmt"
	"gintest/models/trip"
	"github.com/gin-gonic/gin"
)

// GetTripList
// @Summary 获取行程列表
// @Schemes
// @Description 获取行程列表
// @Tags 行程
// @Accept json
// @Produce json
// @Success 200 {string} json "{"code": 200, "msg":"", "data":""}"
// @Router /trip/list [get]
func GetTripList(c *gin.Context) {
	fmt.Println("this is a get trip list action")
	trip.GetTripList()

	data := "123"

	c.JSON(200, gin.H{"code": 200, "data": data, "msg": "成功"})

}
