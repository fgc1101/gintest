package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
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

			var trip Trip
			err := context.ShouldBindJSON(&trip)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(trip)
		}) // 添加行程
		tripRouter.GET("/list", func(context *gin.Context) {
			var list []Trip

			for i := 0; i < 3; i++ {

				trip := Trip{
					Uid:  int32(i) + 1,
					Mode: "来程",
				}

				list = append(list, trip)
			}
			resp := Response{
				Code:    1,
				Message: "Success",
				Data:    list,
			}
			context.JSON(200, &resp)
		}) // 行程列表
		tripRouter.GET("/index", func(context *gin.Context) {

			list := GetTripList()
			context.HTML(http.StatusOK, "index.html", gin.H{
				"data": list,
			})
		})
	}

	err := r.Run(":8009")
	if err != nil {
		fmt.Println("服务器启动发生了错误……")
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}

type Trip struct {
	Uid  int32  `json:"uid" binding:"required"`
	Mode string `json:"mode"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetTripList() []Trip {
	var list []Trip

	for i := 0; i < 3; i++ {

		trip := Trip{
			Uid:  int32(i) + 1,
			Mode: "来程",
		}

		list = append(list, trip)
	}
	return list
}
