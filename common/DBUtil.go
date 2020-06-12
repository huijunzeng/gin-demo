package common

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 定义一个全局对象DB
var DB *gorm.DB

// 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:zeng@19940125...@tcp(http://129.211.34.120:3306)/gin_demo?charset=utf8"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	return DB.DB().Ping()
}

func Close() {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
