package Controller

import (
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": [0]int{},
	})
}
