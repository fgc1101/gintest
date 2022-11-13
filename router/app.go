package router

import (
	"fmt"
	"gintest/controllers/trip"
	_ "gintest/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 返回HTML页面
	r.LoadHTMLGlob("./template/html/*")

	r.Static("/img", "./template/static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// trip 路由组
	tripRouter := r.Group("/trip")
	{
		tripRouter.POST("/add", func(context *gin.Context) {
			// TODO
			fmt.Println(1)
		}) // 添加行程
		tripRouter.GET("/list", trip.GetTripList)  // 行程列表
		tripRouter.GET("/index", trip.GetTripList) // 行程列表

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
