package conf

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"os"
)

var (
	Conf *Config = new(Config)
)

type Config struct {
	Token string `toml:"token"`
	Host  string `toml:"host"`
	Port  string `toml:"port"`
}

func LoadConfig() {
	_, err := toml.DecodeFile("./conf/production.toml", Conf)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
}
