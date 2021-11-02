package Controller

import (
	"fmt"
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
	resp(ctx, 200, msg, data...)
}

// resp 返回
func resp(ctx *gin.Context, code int, msg string, data ...interface{}) {
	resp := Resp{
		Data: data,
	}

	fmt.Println(data)
	if resp.Data == nil {
		fmt.Print("111")
		resp.Data = [0]int{}
	}

	if len(data) == 1 {
		fmt.Print("222")
		resp.Data = data[0]
	}
	if len(data) == 0 || resp.Data == nil {
		fmt.Print("333")
		resp.Data = [0]int{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    msg,
		"data":   resp.Data,
	})
	return
}

// 全局中间件回调函数
func MiddleWare() gin.HandlerFunc {
	// 当客户端有请求来之后, 先执行这个函数
	return func(c *gin.Context) {
		fmt.Println("MiddleWare: 中间件开始执行")
		c.Set("request", "中间件") // 在gin.Context中设置一个值(这个与中间件无关, 只是为了演示)
		// 执行到这个函数时, 会跳转到main函数中, 执行客户端对应的请求回调函数
		c.Next()
		// 执行完对应的回调函数之后, 继续回到这个地方进行执行(但是响应还没有返回给客户端)
		// 获取相应状态码, 并打印相关信息
		status := c.Writer.Status()
		fmt.Println("MiddleWare: 中间件执行结束, status: ", status)
		return
	}
	// 当中间件执行完之后, 才真正把响应返回给客户端
}
