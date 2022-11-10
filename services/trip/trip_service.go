package trip

import (
	"fmt"
	"gintest/models/trip"
	"github.com/gin-gonic/gin"
)

func GetTripList(c *gin.Context) {
	fmt.Println("this is a get trip list action")
	trip.GetTripList()
}
