package services

import (
	"fmt"
	"gin-demo/models"
	mapper "gin-demo/repositories"
)

func GetUserListByPage() []*models.User {
	fmt.Println("==================-------------------------===")
	userList := mapper.GetUserListByPage()
	return userList
}
