package repositories

import (
	"gin-demo/common"
	"gin-demo/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserListByPage() []*models.User {
	var userList []*models.User
	common.DB.Table("t_users").Find(&userList)

	return userList
}
