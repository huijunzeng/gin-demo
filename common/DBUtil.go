package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 定义一个全局对象db
var MysqlDB *gorm.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:zeng@19940125...@tcp(http://129.211.34.120:3306)/gin_demo?charset=utf8"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	MysqlDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return nil
}
