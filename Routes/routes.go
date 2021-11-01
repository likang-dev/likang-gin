package Routes

import (
	"github.com/gin-gonic/gin"
	"likang-project/Controller"
	_ "likang-project/Controller"
)

func Load(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// v1版本接口
	v1 := r.Group("/v1")
	{
		// 用户相关
		user := v1.Group("/user")
		{
			// 登陆接口
			user.POST("/login", Controller.Login)
		}
	}

	// 默认控制器，路由不存在时候访问
	r.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"msg": "404 not found",
		})
	})

}