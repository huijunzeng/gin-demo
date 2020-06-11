package repositories

import (
	db "gin-demo/common"
	models1 "gin-demo/models"
	_ "github.com/go-sql-driver/mysql"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (r UserDao) GetUserListByPage() []models1.User {
	var userList []models1.User
	db.MysqlDB.Table("t_user").Find(&userList)

	return userList
}
