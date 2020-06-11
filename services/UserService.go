package services

import (
	"gin-demo/models"
	"gin-demo/repositories"
)

type UserService struct {
	userDao *repositories.UserDao
}

func NewUserService() *UserService {
	return &UserService{&repositories.UserDao{}}
}

func (u UserService) GetUserListByPage() []models.User {
	userList := u.userDao.GetUserListByPage()
	return userList
}
