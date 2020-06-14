package models

/**
*结构体中的字段名，如果首字母小写的话，则该字段无法被外部包访问和解析，比如，json解析
*user实体类
 */
type User struct {
	Id       int32
	Name     string
	Password string
}
