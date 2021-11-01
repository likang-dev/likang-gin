package Controller

import (
	"github.com/gin-gonic/gin"
)

var data struct{}

func Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
