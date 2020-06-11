package repositories

import (
	"fmt"
	"gin-demo/common"
	models1 "gin-demo/models"
	_ "github.com/go-sql-driver/mysql"
)

func GetUserListByPage() []*models1.User {
	var userList []*models1.User
	fmt.Println("========111111111111==========-------------------------===")
	//db.Db.Where("name=?", "hanyun").Find(&&userList)
	common.InitDB().Table("t_users").Find(&userList)
	fmt.Println("========111111111111=5555555555555555---------------===")
	return userList
}
