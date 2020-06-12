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
	dsn := "root:zeng@19940125...@tcp(http://129.211.34.120:3306)/gin_demo?charset=utf8&parseTime=True&loc=Local"
	// gorm默认表名是结构体名的复数形式（即默认结构体user对应表users），这里设置全局禁用表名复数（即结构体user对应表user，结构体users对应表users）
	//db.SingularTable(true)
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
