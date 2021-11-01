package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Login(context *gin.Context) {
	userInfo := UserInfo{}
	context.BindJSON(&userInfo)
	fmt.Println(userInfo.UserName)
	fmt.Println(userInfo.Password)
	context.JSON(200, userInfo)
}
