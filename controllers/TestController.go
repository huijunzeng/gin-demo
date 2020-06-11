package controllers

import (
	"fmt"
	"gin-demo/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{services.NewUserService()}
}

// @Summary 查询用户列表接口
// @Description 查询用户列表接口
// @Tags 用户信息
// @Success 200 {string} json "{"message":"success"}"
// @Router /api/v1/getUserListByPage [get]
func (u UserController) GetUserListByPage(c *gin.Context) {
	userList := u.userService.GetUserListByPage()
	fmt.Println(userList)
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}
