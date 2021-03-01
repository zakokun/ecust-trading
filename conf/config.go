package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
)

// 配置文件
var c *Config

type Config struct {
	DB    *DBConf
	Trade *TradeConf
}

type TradeConf struct {
	Symbol string
}

type DBConf struct {
	Addr     string
	Port     int64
	User     string
	Password string
	DBName   string
}

func init() {
	configFile := flag.String("c", "./conf/config_template.toml", "conf")
	flag.Parse()

	c = &Config{}
	_, err := toml.DecodeFile(*configFile, c)
	if err != nil {
		fmt.Println("failed to decode config file", configFile, err)
		panic(err)
	}
}
func Get() *Config {
	return c
}
