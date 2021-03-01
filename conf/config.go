package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
)

// 配置文件
var c *Config

type Config struct {
	DB     *DB
	Symbol string
}

type DB struct {
	Addr     string
	Port     int64
	User     string
	Password string
	DBName   string
}

func init() {
	configFile := flag.String("c", "./conf/config_template.toml", "conf")
	flag.Parse()

	Conf := &Config{}
	_, err := toml.DecodeFile(*configFile, Conf)
	if err != nil {
		fmt.Println("failed to decode config file", configFile, err)
		panic(err)
	}
}
func Get() *Config {
	return c
}
