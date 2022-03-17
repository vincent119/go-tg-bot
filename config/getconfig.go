package config

import (
	//"log"
	"os"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
	"fmt"
)

var (
	cfg = pflag.StringP("Config","c","","./config/")
)

type appConf struct {
	Port string `mapstructure:"port"`
}
type Url1 struct {
	LogName string `mapstructure:"logname"`
}
type RedisConfig struct {
	Port string `mapstructure:"port"`
	Auth string `mapstructure:"auth"`
	Host string `mapstructure:"ip"`
	TTL int `mapstructure:"ttl"`
}

type config struct {
	App appConf `mapstructure:"app"`
	Uri1 Url1 `mapstructure:"url1"`
	RedisCfg RedisConfig `mapstructure:"redis"`
}

type LogrusConfig struct {
	Path         string `ini:"path"`
	Level        string `ini:"level"`
	Formatter    string `ini:"formatter"`
	OutputType   string `ini:"output_type"`
	ReportCaller bool   `ini:"report_caller"`
	Suffix       string `ini:"suffix_format"`
}

var cf *viper.Viper
var Conf config
func Init(){
	//log.Print("This is the environment: ", env)
	pflag.Parse()
	cf = viper.New()
	confPath ,_ := os.Getwd()
	if *cfg != ""{
	  cf.SetConfigFile(*cfg)
	} else {
	  cf.AddConfigPath(".")
	  cf.AddConfigPath(confPath)
	}
	cf.SetConfigType("yaml")
	cf.AutomaticEnv()
	err := cf.ReadInConfig()
	if err != nil {
	  panic(err)
	}
	cf.WatchConfig()
	cf.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := cf.Unmarshal(&Conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := cf.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}
	if err := cf.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic(" Config file not found; ignore error if desired")
		} else {
				panic("Config file was found but another error was produced")
		}
  }

}

