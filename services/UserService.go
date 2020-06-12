package services

import (
	"gin-demo/models"
	"gin-demo/repositories"
)

func GetUserListByPage() []*models.User {
	userList := repositories.GetUserListByPage()
	return userList
}
