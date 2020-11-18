package config

/**
	结构体中的字段名，如果首字母小写的话，则该字段无法被外部包访问和解析，比如，json解析
`mapstructure:"username"`
*/
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
	Dbname   string `mapstructure:"dbname"`
	Config   string `mapstructure:"config"`
}

type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
}
