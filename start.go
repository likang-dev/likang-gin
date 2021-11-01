package main

import (
	"github.com/gin-gonic/gin"
	"likang-project/Routes"
)

func main() {
	r := gin.Default()
	Routes.Load(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
