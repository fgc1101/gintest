package main

import (
	"fmt"
	"gintest/router"
)

func main() {
	r := router.Router()

	err := r.Run(":8009")
	if err != nil {
		fmt.Println("服务器启动发生了错误……")
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
