package controllers

import (
	"fmt"
	"gin-demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 查询用户列表接口
// @Description 查询用户列表接口
// @Tags 用户信息
// @Success 200 {string} json "{"message":"success"}"
// @Router /users [get]
func GetUserListByPage(c *gin.Context) {
	fmt.Println("===========================")
	userList := services.GetUserListByPage()
	fmt.Println("=============", len(userList))
	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})
}
