package config


type AppConf struct {
	Port string `mapstructure:"port"`
}



func GetAppConf() *AppConf {
	return 
}