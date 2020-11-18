package models

/**
结构体中的字段名，如果首字母小写的话，则该字段无法被外部包访问和解析，比如，json解析
User实体类
*/
type User struct {
	Id       uint   `json:"id" gorm:"comment:用户ID"`
	Name     string `json:"name" gorm:"comment:用户名"`
	Password string `json:"password" gorm:"comment:用户密码"`
}
