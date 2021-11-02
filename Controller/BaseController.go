package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Resp 返回
type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

// ErrorResp 错误返回值
func ErrorResp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp(ctx, code, msg, data...)
}

// SuccessResp 正确返回值
func SuccessResp(ctx *gin.Context, msg string, data ...interface{}) {
	resp(ctx, 0, msg, data...)
}

// resp 返回
func resp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp := Resp{
		Data: data,
	}
	if len(data) == 1 {
		resp.Data = data[0]
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    msg,
		"data":   resp.Data,
	})
}
