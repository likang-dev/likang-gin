package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func Login(context *gin.Context) {
	userInfo := UserInfo{}
	context.BindJSON(&userInfo)
	fmt.Println(userInfo.UserName)
	fmt.Println(userInfo.Password)
	var user1, user2 UserInfo
	user1.UserName = "zhangsan"
	user1.Password = "123456"
	user1.Address = "北京"
	user2.UserName = "lisi"
	user2.Password = "1234567"
	arr := [...]UserInfo{user1, user2, userInfo}
	SuccessResp(context, "success", arr)
	return
}
