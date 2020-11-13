package config

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
